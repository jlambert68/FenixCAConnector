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
		fenixExecutionConnectorObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend GrpcServer Server
		fenixExecutionConnectorObject.GrpcServer.StopGrpcServer()

	}
}

func fenixExecutionConnectorMain() {

	// Set up BackendObject
	fenixExecutionConnectorObject = &fenixExecutionConnectorObjectStruct{
		logger:     nil,
		GrpcServer: &gRPCServer.FenixExecutionConnectorGrpcObjectStruct{},
	}

	// Init logger
	fenixExecutionConnectorObject.InitLogger("")

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Initiate Logger for gRPC-server
	fenixExecutionConnectorObject.GrpcServer.InitiateLogger(fenixExecutionConnectorObject.logger)

	// Start Backend GrpcServer-server
	fenixExecutionConnectorObject.GrpcServer.InitGrpcServer(fenixExecutionConnectorObject.logger)

}
