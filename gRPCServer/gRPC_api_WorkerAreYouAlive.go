package gRPCServer

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/messagesToExecutionWorkerServer"
	"fmt"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// AreYouAlive - *********************************************************************
//Anyone can check if Fenix Execution Worker server is alive with this service, should be used to check serves for Connector
func (s *fenixExecutionConnectorGrpcServicesServer) WorkerAreYouAlive(ctx context.Context, emptyParameter *fenixExecutionConnectorGrpcApi.EmptyParameter) (*fenixExecutionConnectorGrpcApi.AckNackResponse, error) {

	s.logger.WithFields(logrus.Fields{
		"id": "5c2d4e0c-904a-41d8-81bc-3123641aa6db",
	}).Debug("Incoming 'gRPCServer - ConnectorAreYouAlive'")

	s.logger.WithFields(logrus.Fields{
		"id": "b9003ecf-b686-429b-b603-261f78e9c787",
	}).Debug("Outgoing 'gRPCServer - ConnectorAreYouAlive'")

	// Current user
	userID := "gRPC-api doesn't support UserId"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(userID, fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(emptyParameter.ProtoFileVersionUsedByCaller))
	if returnMessage != nil {

		// Exiting
		return returnMessage, nil
	}

	// Set up instance to use for execution gPRC
	var fenixExecutionWorkerObject *messagesToExecutionWorkerServer.MessagesToExecutionWorkerObjectStruct
	fenixExecutionWorkerObject = &messagesToExecutionWorkerServer.MessagesToExecutionWorkerObjectStruct{
		Logger: s.logger,
		//GcpAccessToken: nil,
	}

	response, responseMessage := fenixExecutionWorkerObject.SendAreYouAliveToFenixExecutionServer()

	// Create Error Codes
	var errorCodes []fenixExecutionConnectorGrpcApi.ErrorCodesEnum

	ackNackResponseMessage := &fenixExecutionConnectorGrpcApi.AckNackResponse{
		AckNack:                         response,
		Comments:                        fmt.Sprintf("The response from Worker is '%s'", responseMessage),
		ErrorCodes:                      errorCodes,
		ProtoFileVersionUsedByConnector: fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(common_config.GetHighestConnectorProtoFileVersion()),
	}

	return ackNackResponseMessage, nil

}
