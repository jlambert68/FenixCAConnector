package main

import (
	"FenixCAConnector/gRPCServer"
	"github.com/sirupsen/logrus"
)

type fenixExecutionConnectorObjectStruct struct {
	logger     *logrus.Logger
	GrpcServer *gRPCServer.FenixExecutionConnectorGrpcObjectStruct
}

// Variable holding everything together
var fenixExecutionConnectorObject *fenixExecutionConnectorObjectStruct
