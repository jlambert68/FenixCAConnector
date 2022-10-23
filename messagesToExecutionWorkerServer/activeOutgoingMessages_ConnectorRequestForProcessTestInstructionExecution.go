package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/gcp"
	"FenixCAConnector/restCallsToCAEngine"
	"context"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
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

				// Send response and start processing TestInstruction in parallell
				go func() {

					// Call 'CA' backend to convert 'TestInstruction' into useful structure later to be used by FangEngine

					var fangEngineRestApiMessageValues *restCallsToCAEngine.FangEngineRestApiMessageStruct
					fangEngineRestApiMessageValues, err = restCallsToCAEngine.ConvertTestInstructionIntoFangEngineRestCallMessage(processTestInstructionExecutionReveredRequest)

					// Generate response depending on if the 'TestInstruction' could be converted into useful FangEngine-information or not
					var processTestInstructionExecutionReversedResponse *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse
					if err != nil {
						// Couldn't convert into FangEngine-messageType
						timeAtDurationEnd := time.Now()

						// Generate response message to Worker, that conversion didn't work out
						processTestInstructionExecutionReversedResponse = &fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse{
							AckNackResponse: &fenixExecutionWorkerGrpcApi.AckNackResponse{
								AckNack:                      false,
								Comments:                     err.Error(),
								ErrorCodes:                   nil,
								ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion()),
							},
							TestInstructionExecutionUuid:   processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionExecutionUuid,
							ExpectedExecutionDuration:      timestamppb.New(timeAtDurationEnd),
							TestInstructionCanBeReExecuted: true,
						}
					} else {
						// Generate duration for Execution:: TODO This is only for test and should be done in another way later
						executionDuration := time.Minute * 5
						timeAtDurationEnd := time.Now().Add(executionDuration)

						// Generate OK response message to Worker
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
					}

					// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
					couldSend, _ := toExecutionWorkerObject.SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer(processTestInstructionExecutionReversedResponse)

					// If response could be sent back to Worker then execute TestInstruction
					if couldSend == true {

						// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
						couldSend, returnMessage := toExecutionWorkerObject.SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer(processTestInstructionExecutionReversedResponse)

						if couldSend == false {
							common_config.Logger.WithFields(logrus.Fields{
								"ID":            "95dddb21-0895-4016-9cb5-97ab4568f30b",
								"returnMessage": returnMessage,
							}).Error("Couldn't send response to Worker")

						}

					} else {

						// Send TestInstruction to FangEngine using RestCall
						var restResponse *http.Response
						restResponse, err = restCallsToCAEngine.PostTestInstructionUsingRestCall(fangEngineRestApiMessageValues)

						// Convert response from restCall into 'Fenix-world-data'
						var testInstructionExecutionStatus fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum
						if err != nil {
							testInstructionExecutionStatus = fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION
						} else {
							switch restResponse.StatusCode {
							case http.StatusOK: // 200
								testInstructionExecutionStatus = fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_FINISHED_OK
							case http.StatusBadRequest: // 400 TODO use correct error
								testInstructionExecutionStatus = fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_FINISHED_NOT_OK

							case http.StatusInternalServerError: // 500
								testInstructionExecutionStatus = fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION

							default:
								// Unhandled response code

								common_config.Logger.WithFields(logrus.Fields{
									"ID":                      "f6d86465-9a3c-4277-9730-929537f1b42b",
									"restResponse.StatusCode": restResponse.StatusCode,
								}).Error("Unhandled response from FangEngine")

								testInstructionExecutionStatus = fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION
							}
						}

						// Generate response message to Worker
						var finalTestInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
						finalTestInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
							ClientSystemIdentification: &fenixExecutionWorkerGrpcApi.ClientSystemIdentificationMessage{
								DomainUuid:                   string(Domains.DomainUUID_CA),
								ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion()),
							},
							TestInstructionExecutionUuid:         processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionExecutionUuid,
							TestInstructionExecutionStatus:       testInstructionExecutionStatus,
							TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
						}

						// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
						couldSend, returnMessage := toExecutionWorkerObject.SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer(finalTestInstructionExecutionResultMessage)

						if couldSend == false {
							common_config.Logger.WithFields(logrus.Fields{
								"ID": "95dddb21-0895-4016-9cb5-97ab4568f30b",
								"finalTestInstructionExecutionResultMessage": finalTestInstructionExecutionResultMessage,
								"returnMessage": returnMessage,
							}).Error("Couldn't send repsonse to Worker")
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
