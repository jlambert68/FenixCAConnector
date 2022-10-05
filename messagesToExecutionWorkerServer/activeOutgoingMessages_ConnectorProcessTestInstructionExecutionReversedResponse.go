package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"context"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer - When a TestInstruction has been fully executed the Client use this to inform the results of the execution result to the Worker (who the forward the message to the Execution Server)
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer(processTestInstructionExecutionReversedResponse *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse) (bool, string) {

	toExecutionWorkerObject.Logger.WithFields(logrus.Fields{
		"id": "7deef335-37fb-462c-978c-5a97a52c207f",
	}).Debug("Incoming 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

	toExecutionWorkerObject.Logger.WithFields(logrus.Fields{
		"id": "f05c825b-16cf-4cc0-8e7a-37e375c24d17",
	}).Debug("Outgoing 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

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
		toExecutionWorkerObject.Logger.WithFields(logrus.Fields{
			"ID": "209c1aaa-b5b6-4d4a-a04c-e3b328ac1eaf",
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

	returnMessage, err := fenixExecutionWorkerGrpcClient.ConnectorProcessTestInstructionExecutionReversedResponse(ctx, processTestInstructionExecutionReversedResponse)

	// Shouldn't happen
	if err != nil {
		toExecutionWorkerObject.Logger.WithFields(logrus.Fields{
			"ID":    "bb37e04d-2154-47df-8eca-ea076a132a59",
			"error": err,
		}).Error("Problem to do gRPC-call to Fenix Execution Worker for 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

		return false, err.Error()

	} else if returnMessage.AckNack == false {
		// FenixTestDataSyncServer couldn't handle gPRC call
		toExecutionWorkerObject.Logger.WithFields(logrus.Fields{
			"ID":                               "e72c61f0-feb4-41d2-a10c-5989bca92cc2",
			"Message from Fenix Worker Server": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to Worker Server for 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

		return false, err.Error()
	}

	toExecutionWorkerObject.Logger.WithFields(logrus.Fields{
		"ID": "b48ae8cc-a145-4527-b417-b3bb815824fc",
		"processTestInstructionExecutionReversedResponse": processTestInstructionExecutionReversedResponse,
	}).Debug("Response regarding that worker received a TestInstruction to execute was successfully sent back to worker")

	return returnMessage.AckNack, returnMessage.Comments

}
