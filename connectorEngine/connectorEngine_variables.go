package connectorEngine

import (
	"FenixCAConnector/messagesToExecutionWorkerServer"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

type TestInstructionExecutionEngineStruct struct {
	logger                                                               *logrus.Logger
	CommandChannelReference                                              *ExecutionEngineChannelType
	messagesToExecutionWorkerObjectReference                             *messagesToExecutionWorkerServer.MessagesToExecutionWorkerObjectStruct
	ongoingTimerOrConnectionForCallingWorkerFortestInstructionsToExecute bool
}

// ExecutionEngineCommandChannel
var ExecutionEngineCommandChannel ExecutionEngineChannelType

type ExecutionEngineChannelType chan ChannelCommandStruct

type ChannelCommandType uint8

const (
	ChannelCommandSendAreYouAliveToFenixWorkerServer ChannelCommandType = iota
	ChannelCommandTriggerReportProcessingCapability
	ChannelCommandTriggerReportCompleteTestInstructionExecutionResult
	ChannelCommandTriggerReportCurrentTestInstructionExecutionResult
	ChannelCommandTriggerSendAllLogPostForExecution
	ChannelCommandTriggerRequestForTestInstructionExecutionToProcess
	ChannelCommandTriggerRequestForTestInstructionExecutionToProcessIn5Minutes
)

type ChannelCommandStruct struct {
	ChannelCommand                                        ChannelCommandType
	ReportCompleteTestInstructionExecutionResultParameter ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct
}

// ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct
// Parameter used when to forward the final execution result for a TestInstruction
type ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct struct {
	TriggerTestInstructionExecutionResultMessage *fenixExecutionConnectorGrpcApi.TriggerTestInstructionExecutionResultMessage
}
