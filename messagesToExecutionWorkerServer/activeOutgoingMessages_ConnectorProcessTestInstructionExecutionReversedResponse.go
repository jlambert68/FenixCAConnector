package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/gcp"
	"context"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer - When a TestInstruction has been fully executed the Client use this to inform the results of the execution result to the Worker (who the forward the message to the Execution Server)
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer(processTestInstructionExecutionReversedResponse *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse) (bool, string) {

	common_config.Logger.WithFields(logrus.Fields{
		"id": "7deef335-37fb-462c-978c-5a97a52c207f",
	}).Debug("Incoming 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

	common_config.Logger.WithFields(logrus.Fields{
		"id": "f05c825b-16cf-4cc0-8e7a-37e375c24d17",
	}).Debug("Outgoing 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string

	ctx = context.Background()

	// Set up connection to Server
	ctx, err := toExecutionWorkerObject.SetConnectionToFenixExecutionWorkerServer(ctx)
	if err != nil {
		return false, err.Error()
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		common_config.Logger.WithFields(logrus.Fields{
			"ID": "209c1aaa-b5b6-4d4a-a04c-e3b328ac1eaf",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP && common_config.GCPAuthentication == true {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			return false, returnMessageString
		}

	}

	returnMessage, err := fenixExecutionWorkerGrpcClient.ConnectorProcessTestInstructionExecutionReversedResponse(ctx, processTestInstructionExecutionReversedResponse)

	// Shouldn't happen
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"ID":    "bb37e04d-2154-47df-8eca-ea076a132a59",
			"error": err,
		}).Error("Problem to do gRPC-call to Fenix Execution Worker for 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

		return false, err.Error()

	} else if returnMessage.AckNack == false {
		// FenixTestDataSyncServer couldn't handle gPRC call
		common_config.Logger.WithFields(logrus.Fields{
			"ID":                        "7763f7d1-9a5e-4407-b97b-0737455c6e54",
			"Message from Fenix Worker": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to Worker for 'SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer'")

		return false, err.Error()
	}

	common_config.Logger.WithFields(logrus.Fields{
		"ID": "b48ae8cc-a145-4527-b417-b3bb815824fc",
		"processTestInstructionExecutionReversedResponse": processTestInstructionExecutionReversedResponse,
	}).Debug("Response regarding that worker received a TestInstruction to execute was successfully sent back to worker")

	return returnMessage.AckNack, returnMessage.Comments

}
