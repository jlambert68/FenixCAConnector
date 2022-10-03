package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"context"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

// InitiateConnectorRequestForProcessTestInstructionExecution
// This gPRC-methods is used when a Execution Connector needs to have its TestInstruction assignments using reverse streaming
// Execution Connector opens the gPRC-channel and assignments are then streamed back to Connector from Worker
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) InitiateConnectorRequestForProcessTestInstructionExecution() {

	var ctx context.Context
	var returnMessageAckNack bool

	// Set up connection to Server
	err := toExecutionWorkerObject.SetConnectionToFenixExecutionWorkerServer()
	if err != nil {
		return
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID": "5f02b94f-b07d-4bd7-9607-89cf712824c9",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP {

		// Add Access token
		ctx, returnMessageAckNack, _ = toExecutionWorkerObject.generateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			return
		}

	}

	// Set up call parameter
	emptyParameter := &fenixExecutionWorkerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion())}

	// Start up stream from Worker server
	stream, err := fenixExecutionWorkerGrpcClient.ConnectorRequestForProcessTestInstructionExecution(ctx, emptyParameter)

	// Couldn't connect to Worker
	if err != nil {
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID":  "d9ab0434-1121-4e2e-95e7-3e1cc99656b0",
			"err": err,
		}).Debug("Couldn't open stream from Worker Server. Will wait 5 minutes and try again")

		return
	}

	// Local channel to decide when Server stopped sending
	done := make(chan bool)

	// Run stream receiver as a go-routine
	go func() {
		for {
			processTestInstructionExecutionReveredRequest, err := stream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				toExecutionWorkerObject.logger.WithFields(logrus.Fields{
					"ID":  "3439f49f-d7d5-477e-9a6b-cfa5ed355bfe",
					"err": err,
				}).Debug("Got some error when receiving TestInstructionExecutionsRequests from Worker, reconnect in 5 minutes")

				done <- true //close(done)
				return

			}

			toExecutionWorkerObject.logger.WithFields(logrus.Fields{
				"ID": "d1ea4370-3e8e-4d2b-9626-a193213e091a",
				"processTestInstructionExecutionReveredRequest": processTestInstructionExecutionReveredRequest,
			}).Debug("Receive TestInstructionExecution from Worker")

			// Call 'CA' backend to execute TestInstruction
			// TODO send TestInstruction over CommandChannel

		}
	}()

	// Server stopped sending so reconnect again in 5 minutes
	<-done
	toExecutionWorkerObject.logger.WithFields(logrus.Fields{
		"ID": "0b5fdb7c-91aa-4dfc-b587-7b6cef83d224",
	}).Debug("Server stopped sending so reconnect again in 5 minutes")

}
