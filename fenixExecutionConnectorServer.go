package main

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/connectorEngine"
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

func fenixExecutionConnectorMain(loggerFileName string) {

	// Set up BackendObject
	fenixExecutionConnectorObject = &fenixExecutionConnectorObjectStruct{
		logger:     nil,
		GrpcServer: &gRPCServer.FenixExecutionConnectorGrpcObjectStruct{},
	}

	// Init logger
	fenixExecutionConnectorObject.InitLogger(loggerFileName)

	common_config.Logger = fenixExecutionConnectorObject.logger

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Initiate CommandChannel
	connectorEngine.ExecutionEngineCommandChannel = make(chan connectorEngine.ChannelCommandStruct)

	// Start ChannelCommand Engine
	fenixExecutionConnectorObject.testInstructionExecutionEngine.CommandChannelReference = &connectorEngine.ExecutionEngineCommandChannel
	fenixExecutionConnectorObject.testInstructionExecutionEngine.InitiateTestInstructionExecutionEngineCommandChannelReader(connectorEngine.ExecutionEngineCommandChannel)

	// Initiate  gRPC-server
	fenixExecutionConnectorObject.GrpcServer.InitiategRPCObject(fenixExecutionConnectorObject.logger)

	// Create Message for CommandChannel to connect to Worker to be able to get TestInstructions to Execute
	triggerTestInstructionExecutionResultMessage := &fenixExecutionConnectorGrpcApi.TriggerTestInstructionExecutionResultMessage{}
	channelCommand := connectorEngine.ChannelCommandStruct{
		ChannelCommand: connectorEngine.ChannelCommandTriggerRequestForTestInstructionExecutionToProcess,
		ReportCompleteTestInstructionExecutionResultParameter: connectorEngine.ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct{
			TriggerTestInstructionExecutionResultMessage: triggerTestInstructionExecutionResultMessage},
	}

	// Send message on channel
	connectorEngine.ExecutionEngineCommandChannel <- channelCommand

	// Start Backend GrpcServer-server
	fenixExecutionConnectorObject.GrpcServer.InitGrpcServer(fenixExecutionConnectorObject.logger)

}
