package gRPCServer

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/connectorEngine"
	"context"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// TriggerReportCompleteTestInstructionExecutionResult
// Trigger Connector to inform Worker of the final execution results for an execution
func (s *fenixExecutionConnectorGrpcServicesServer) TriggerReportCompleteTestInstructionExecutionResult(ctx context.Context, triggerTestInstructionExecutionResultMessage *fenixExecutionConnectorGrpcApi.TriggerTestInstructionExecutionResultMessage) (ackNackResponse *fenixExecutionConnectorGrpcApi.AckNackResponse, err error) {

	s.logger.WithFields(logrus.Fields{
		"id": "a6acfc70-deb0-42d6-b6f9-40c3df66256a",
	}).Debug("Incoming 'gRPCServer - TriggerReportCompleteTestInstructionExecutionResult'")

	defer s.logger.WithFields(logrus.Fields{
		"id": "0778c0a5-71ee-4b9a-b9bc-f7fc8fecc93d",
	}).Debug("Outgoing 'gRPCServer - TriggerReportCompleteTestInstructionExecutionResult'")

	// Calling system
	userId := "External Trigger"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(userId, triggerTestInstructionExecutionResultMessage.ProtoFileVersionUsedByCaller)
	if returnMessage != nil {

		return returnMessage, nil
	}

	// Send Message on CommandChannel to be able to send Result back to Fenix Execution Server
	channelCommand := connectorEngine.ChannelCommandStruct{
		ChannelCommand: connectorEngine.ChannelCommandTriggerReportCompleteTestInstructionExecutionResult,
		ReportCompleteTestInstructionExecutionResultParameter: connectorEngine.ChannelCommandSendReportCompleteTestInstructionExecutionResultToFenixExecutionServerStruct{
			TriggerTestInstructionExecutionResultMessage: triggerTestInstructionExecutionResultMessage},
	}

	*s.CommandChannelReference <- channelCommand

	// Generate response
	ackNackResponse = &fenixExecutionConnectorGrpcApi.AckNackResponse{
		AckNack:                         true,
		Comments:                        "",
		ErrorCodes:                      nil,
		ProtoFileVersionUsedByConnector: fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(common_config.GetHighestConnectorProtoFileVersion()),
	}

	return ackNackResponse, nil

}
