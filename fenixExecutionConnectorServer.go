package main

import (
	"FenixCAConnector/gRPCServer"
	"github.com/sirupsen/logrus"
)

// Used for only process cleanup once
var cleanupProcessed = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		FenixExecutionConnectorObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend GrpcServer Server
		FenixExecutionConnectorObject.GrpcServer.StopGrpcServer()

	}
}

func fenixExecutionConnectorMain() {

	// Set up BackendObject
	FenixExecutionConnectorObject = &fenixExecutionConnectorObjectStruct{
		logger:     nil,
		GrpcServer: &gRPCServer.FenixExecutionConnectorGrpcObjectStruct{},
	}

	// Init logger
	FenixExecutionConnectorObject.InitLogger("")

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Initiate Logger for gRPC-server
	FenixExecutionConnectorObject.GrpcServer.InitiateLogger(FenixExecutionConnectorObject.logger)

	// Start Backend GrpcServer-server
	FenixExecutionConnectorObject.GrpcServer.InitGrpcServer(FenixExecutionConnectorObject.logger)

}
