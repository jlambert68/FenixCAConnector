package main

import (
	"FenixCAConnector/gRPCServer"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
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
	go fenixExecutionConnectorObject.GrpcServer.InitGrpcServer(fenixExecutionConnectorObject.logger)

	// Call

	// Create Message for CommandChannel to connect to Worker to be able to get TestInstructions to Execute
	triggerTestInstructionExecutionResultMessage := &fenixExecutionConnectorGrpcApi.TriggerTestInstructionExecutionResultMessage{}
	channelCommand := ChannelCommandStruct{
		ChannelCommand: ChannelCommandTriggerRequestForTestInstructionExecutionToProcessIn5Minutes,
		ReportCompleteTestInstructionExecutionResultParameter: ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct{
			TriggerTestInstructionExecutionResultMessage: triggerTestInstructionExecutionResultMessage},
	}

	// Send message on channel
	*executionEngine.CommandChannelReference <- channelCommand

}
