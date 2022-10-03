package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"context"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendAreYouAliveToFenixExecutionServer - Ask Fenix Execution Server to check if it's up and running
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) SendAreYouAliveToFenixExecutionServer() (bool, string) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string

	// Set up connection to Server
	err := toExecutionWorkerObject.SetConnectionToFenixExecutionWorkerServer()
	if err != nil {
		return false, err.Error()
	}

	// Create the message with all test data to be sent to Fenix
	emptyParameter := &fenixExecutionWorkerGrpcApi.EmptyParameter{

		ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion()),
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID": "c5ba19bd-75ff-4366-818d-745d4d7f1a52",
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

	returnMessage, err := fenixExecutionWorkerGrpcClient.ConnectorAreYouAlive(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID":    "818aaf0b-4112-4be4-97b9-21cc084c7b8b",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixExecutionServer for 'SendAreYouAliveToFenixExecutionServer'")

		return false, err.Error()

	} else if returnMessage.AckNack == false {
		// FenixTestDataSyncServer couldn't handle gPRC call
		toExecutionWorkerObject.logger.WithFields(logrus.Fields{
			"ID":                                  "2ecbc800-2fb6-4e88-858d-a421b61c5529",
			"Message from Fenix Execution Server": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixExecutionServer for 'SendAreYouAliveToFenixExecutionServer'")

		return false, err.Error()
	}

	return returnMessage.AckNack, returnMessage.Comments

}
