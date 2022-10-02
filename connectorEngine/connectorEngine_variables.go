package connectorEngine

import (
	"FenixExecutionWorker/messagesToExecutionWorkerServer"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

type TestInstructionExecutionEngineStruct struct {
	logger                                   *logrus.Logger
	CommandChannelReference                  *ExecutionEngineChannelType
	messagesToExecutionWorkerObjectReference *messagesToExecutionWorkerServer.MessagesToExecutionWorkerObjectStruct
}

// ExecutionEngineCommandChannel
var ExecutionEngineCommandChannel ExecutionEngineChannelType

type ExecutionEngineChannelType chan ChannelCommandStruct

type ChannelCommandType uint8

const (
	ChannelCommandSendAreYouAliveToFenixExecutionServer ChannelCommandType = iota
	ChannelCommandSendReportProcessingCapabilityToFenixExecutionServer
	ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServer
)

type ChannelCommandStruct struct {
	ChannelCommand                                        ChannelCommandType
	ReportCompleteTestInstructionExecutionResultParameter ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct
}

// ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct
// Parameter used when to forward the final execution result for a TestInstruction
type ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct struct {
	FinalTestInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
}
