package main

import (
	"FenixCAConnector/connectorEngine"
	"FenixCAConnector/gRPCServer"
	"github.com/sirupsen/logrus"
)

type fenixExecutionConnectorObjectStruct struct {
	logger                  *logrus.Logger
	GrpcServer              *gRPCServer.FenixExecutionConnectorGrpcObjectStruct
	CommandChannelReference *connectorEngine.ExecutionEngineChannelType
}

// Variable holding everything together
var fenixExecutionConnectorObject *fenixExecutionConnectorObjectStruct
