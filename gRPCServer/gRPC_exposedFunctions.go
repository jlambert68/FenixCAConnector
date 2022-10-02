package gRPCServer

import "github.com/sirupsen/logrus"

// InitiateLogger - Initiate local logger object
func (fenixExecutionConnectorGrpcObject *FenixExecutionConnectorGrpcObjectStruct) InitiateLogger(logger *logrus.Logger) {

	fenixExecutionConnectorGrpcObject.logger = logger
}

// InitiateLocalObject - Initiate local 'ExecutionConnectorGrpcObject'
func (fenixExecutionConnectorGrpcObject *FenixExecutionConnectorGrpcObjectStruct) InitiateLocalObject(inFenixExecutionConnectorGrpcObject *FenixExecutionConnectorGrpcObjectStruct) {

	fenixExecutionConnectorGrpcObject.ExecutionConnectorGrpcObject = inFenixExecutionConnectorGrpcObject
}
