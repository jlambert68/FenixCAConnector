package gRPCServer

import (
	"FenixCAConnector/common_config"
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

	ackNackResponse := &fenixExecutionConnectorGrpcApi.AckNackResponse{
		AckNack:                         true,
		Comments:                        "I'am alive.",
		ErrorCodes:                      nil,
		ProtoFileVersionUsedByConnector: fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(common_config.GetHighestConnectorProtoFileVersion()),
	}

	return ackNackResponse, nil
}
