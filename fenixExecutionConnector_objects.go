package main

import (
	"FenixExecutionWorker/gRPCServer"
	"github.com/sirupsen/logrus"
)

type fenixExecutionConnectorObjectStruct struct {
	logger     *logrus.Logger
	GrpcServer *gRPCServer.FenixExecutionConnectorGrpcObjectStruct
}

// Variable holding everything together
var FenixExecutionConnectorObject *fenixExecutionConnectorObjectStruct
