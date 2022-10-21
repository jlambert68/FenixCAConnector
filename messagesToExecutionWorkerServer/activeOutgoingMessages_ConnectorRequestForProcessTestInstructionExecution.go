package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/gcp"
	"FenixCAConnector/restCallsToCAEngine"
	"context"
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"time"
)

// InitiateConnectorRequestForProcessTestInstructionExecution
// This gPRC-methods is used when a Execution Connector needs to have its TestInstruction assignments using reverse streaming
// Execution Connector opens the gPRC-channel and assignments are then streamed back to Connector from Worker
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) InitiateConnectorRequestForProcessTestInstructionExecution() {

	common_config.Logger.WithFields(logrus.Fields{
		"id": "c8e7cbdb-46bd-4545-a472-056fff940365",
	}).Debug("Incoming 'InitiateConnectorRequestForProcessTestInstructionExecution'")

	common_config.Logger.WithFields(logrus.Fields{
		"id": "be16c2a2-4443-4e55-8ad1-9c8478a75e12",
	}).Debug("Outgoing 'InitiateConnectorRequestForProcessTestInstructionExecution'")

	var ctx context.Context
	var returnMessageAckNack bool

	ctx = context.Background()

	// Set up connection to Server
	ctx, err := toExecutionWorkerObject.SetConnectionToFenixExecutionWorkerServer(ctx)
	if err != nil {
		return
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background()) //, 30*time.Second)
	defer func() {
		common_config.Logger.WithFields(logrus.Fields{
			"ID": "5f02b94f-b07d-4bd7-9607-89cf712824c9",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP && common_config.GCPAuthentication == true {

		// Add Access token
		ctx, returnMessageAckNack, _ = gcp.Gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			return
		}

	}

	// Set up call parameter
	emptyParameter := &fenixExecutionWorkerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion())}

	// Start up streamClient from Worker server
	streamClient, err := fenixExecutionWorkerGrpcClient.ConnectorRequestForProcessTestInstructionExecution(ctx, emptyParameter)

	// Couldn't connect to Worker
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"ID":  "d9ab0434-1121-4e2e-95e7-3e1cc99656b0",
			"err": err,
		}).Error("Couldn't open streamClient from Worker Server. Will wait 5 minutes and try again")

		return
	}

	// Local channel to decide when Server stopped sending
	done := make(chan bool)

	// Run streamClient receiver as a go-routine
	go func() {
		for {
			processTestInstructionExecutionReveredRequest, err := streamClient.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				common_config.Logger.WithFields(logrus.Fields{
					"ID":  "3439f49f-d7d5-477e-9a6b-cfa5ed355bfe",
					"err": err,
				}).Error("Got some error when receiving TestInstructionExecutionsRequests from Worker, reconnect in 5 minutes")

				done <- true //close(done)
				return

			}

			// Check if message counts as a "keep Alive message, message is 'nil
			if processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionName == "KeepAlive" {
				// Is a keep alive message
				common_config.Logger.WithFields(logrus.Fields{
					"ID": "08b86c8d-81ba-4664-8cb5-8e53140dc870",
					"processTestInstructionExecutionReveredRequest": processTestInstructionExecutionReveredRequest,
				}).Debug("'Keep alive' message received from Worker")

			} else {
				// Is a standard TestInstruction to execute by Connector backend
				common_config.Logger.WithFields(logrus.Fields{
					"ID": "d1ea4370-3e8e-4d2b-9626-a193213e091a",
					"processTestInstructionExecutionReveredRequest": processTestInstructionExecutionReveredRequest,
				}).Debug("Receive TestInstructionExecution from Worker")

				// Generate duration for Execution:: TODO This is only for test and should be done in another way later
				executionDuration := time.Minute * 5
				timeAtDurationEnd := time.Now().Add(executionDuration)

				// Generate response message to Worker
				var processTestInstructionExecutionReversedResponse *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse
				processTestInstructionExecutionReversedResponse = &fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse{
					AckNackResponse: &fenixExecutionWorkerGrpcApi.AckNackResponse{
						AckNack:                      true,
						Comments:                     "",
						ErrorCodes:                   nil,
						ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion()),
					},
					TestInstructionExecutionUuid:   processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionExecutionUuid,
					ExpectedExecutionDuration:      timestamppb.New(timeAtDurationEnd),
					TestInstructionCanBeReExecuted: false,
				}

				// Send response and start processing TestInstruction in parallell
				go func() {
					// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
					couldSend, _ := toExecutionWorkerObject.SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer(processTestInstructionExecutionReversedResponse)

					// If response could be sent back to Worker then execute TestInstruction
					if couldSend == true {

						// Call 'CA' backend to execute TestInstruction
						fmt.Println("Execution TestInstruction at Custody Arrangement-Automation")
						err = restCallsToCAEngine.ConvertTestInstructionIntoFangEngineRestCallMessage(processTestInstructionExecutionReveredRequest)

						if err != nil {
							// Couldn't convert into FangEngine-messageType
							//TODO Send response about failed TestInstruction to Worker
						}

					}
				}()

			}

		}
	}()

	// Server stopped sending so reconnect again in 5 minutes
	<-done
	common_config.Logger.WithFields(logrus.Fields{
		"ID": "0b5fdb7c-91aa-4dfc-b587-7b6cef83d224",
	}).Debug("Server stopped sending so reconnect again in 5 minutes")

}
