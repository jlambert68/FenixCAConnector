package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"context"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer - When a TestInstruction has been fully executed the Client use this to inform the results of the execution result to the Worker (who the forward the message to the Execution Server)
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer(finalTestInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage) (bool, string) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string

	// Set up connection to Server
	err := toExecutionWorkerObject.SetConnectionToFenixExecutionWorkerServer()
	if err != nil {
		return false, err.Error()
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
		ctx, returnMessageAckNack, returnMessageString = toExecutionWorkerObject.generateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			return false, returnMessageString
		}

	}

	returnMessage, err := fenixExecutionWorkerGrpcClient.ConnectorReportCompleteTestInstructionExecutionResult(ctx, finalTestInstructionExecutionResultMessage)

	// Shouldn't happen
	if err != nil {
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID":    "ebe601e0-14b9-42c5-8f8f-960acec80433",
			"error": err,
		}).Error("Problem to do gRPC-call to Fenix Execution Worker for 'SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer'")

		return false, err.Error()

	} else if returnMessage.AckNack == false {
		// FenixTestDataSyncServer couldn't handle gPRC call
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID":                                  "e72c61f0-feb4-41d2-a10c-5989bca92cc2",
			"Message from Fenix Execution Server": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixExecutionServer for 'SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer'")

		return false, err.Error()
	}

	return returnMessage.AckNack, returnMessage.Comments

}
