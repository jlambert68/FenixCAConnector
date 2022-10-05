package gRPCServer

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/messagesToExecutionWorkerServer"
	"FenixCAConnector/systemSpecific_CA"
	"context"
	"fmt"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// TriggerReportCompleteTestInstructionExecutionResult
// Trigger Connector to inform Worker of the final execution results for an execution
func (s *fenixExecutionConnectorGrpcServicesServer) TriggerReportCompleteTestInstructionExecutionResult(ctx context.Context, triggerTestInstructionExecutionResultMessage *fenixExecutionConnectorGrpcApi.TriggerTestInstructionExecutionResultMessage) (ackNackResponse *fenixExecutionConnectorGrpcApi.AckNackResponse, err error) {

	s.logger.WithFields(logrus.Fields{
		"id": "314111b4-6160-4662-a40a-91643e6d1fda",
	}).Debug("Incoming 'gRPCServer - TriggerReportCompleteTestInstructionExecutionResult'")

	defer s.logger.WithFields(logrus.Fields{
		"id": "b580c24b-3f2a-471e-ad0a-abb1cedf8619",
	}).Debug("Outgoing 'gRPCServer - TriggerReportCompleteTestInstructionExecutionResult'")

	// Current user
	userID := "gRPC-api doesn't support UserId"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(userID, fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(triggerTestInstructionExecutionResultMessage.ProtoFileVersionUsedByCaller))
	if returnMessage != nil {

		// Exiting
		return returnMessage, nil
	}

	// Set up instance to use for execution gPRC
	var fenixExecutionWorkerObject *messagesToExecutionWorkerServer.MessagesToExecutionWorkerObjectStruct
	fenixExecutionWorkerObject = &messagesToExecutionWorkerServer.MessagesToExecutionWorkerObjectStruct{Logger: s.logger}

	// Create 'FinalTestInstructionExecutionResultMessage'
	var finalTestInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
	finalTestInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
		ClientSystemIdentification: &fenixExecutionWorkerGrpcApi.ClientSystemIdentificationMessage{
			DomainUuid:                   systemSpecific_CA.DomainUuid,
			ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion()),
		},
		TestInstructionExecutionUuid:   triggerTestInstructionExecutionResultMessage.TestInstructionExecutionUuid,
		TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum(triggerTestInstructionExecutionResultMessage.TestInstructionExecutionStatus),
	}

	succeededToSend, responseMessage := fenixExecutionWorkerObject.SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer(finalTestInstructionExecutionResultMessage)

	// Create Error Codes
	var errorCodes []fenixExecutionConnectorGrpcApi.ErrorCodesEnum

	ackNackResponseMessage := &fenixExecutionConnectorGrpcApi.AckNackResponse{
		AckNack:                         succeededToSend,
		Comments:                        fmt.Sprintf("The response from Worker is '%s'", responseMessage),
		ErrorCodes:                      errorCodes,
		ProtoFileVersionUsedByConnector: fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(common_config.GetHighestConnectorProtoFileVersion()),
	}

	return ackNackResponseMessage, nil

}
