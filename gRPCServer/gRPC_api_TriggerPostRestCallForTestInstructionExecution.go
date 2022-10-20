package gRPCServer

import (
	"FenixCAConnector/common_config"
	"context"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// TriggerSendAllLogPostForExecution
// Trigger Connector to inform Worker of all log posts that have been produced for an execution
func (s *fenixExecutionConnectorGrpcServicesServer) TriggerPostRestCallForTestInstructionExecution(ctx context.Context, triggerTestInstructionExecutionResultMessage *fenixExecutionConnectorGrpcApi.TriggerTestInstructionExecutionResultMessage) (ackNackResponse *fenixExecutionConnectorGrpcApi.AckNackResponse, err error) {

	s.logger.WithFields(logrus.Fields{
		"id": "fca33679-705d-47a7-8601-18abaa8be1a4",
	}).Debug("Incoming 'gRPCServer - TriggerPostRestCallForTestInstructionExecution'")

	defer s.logger.WithFields(logrus.Fields{
		"id": "9770a757-83d9-4120-adb5-cfe09875c6ba",
	}).Debug("Outgoing 'gRPCServer - TriggerPostRestCallForTestInstructionExecution'")

	// Calling system
	userId := "External Trigger"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(userId, triggerTestInstructionExecutionResultMessage.ProtoFileVersionUsedByCaller)
	if returnMessage != nil {

		return returnMessage, nil
	}

	//TODO Send RestCall

	// Generate response
	ackNackResponse = &fenixExecutionConnectorGrpcApi.AckNackResponse{
		AckNack:                         true,
		Comments:                        "",
		ErrorCodes:                      nil,
		ProtoFileVersionUsedByConnector: fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(common_config.GetHighestConnectorProtoFileVersion()),
	}

	return ackNackResponse, nil

}
