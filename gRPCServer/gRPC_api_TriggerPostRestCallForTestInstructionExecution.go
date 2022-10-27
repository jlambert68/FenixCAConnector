package gRPCServer

import (
	"FenixCAConnector/common_config"
	"context"
	fenixExecutionConnectorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionConnectorGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	_ "github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TriggerSendAllLogPostForExecution
// Trigger Connector to inform Worker of all log posts that have been produced for an execution
func (s *fenixExecutionConnectorGrpcServicesServer) TriggerPostRestCallForTestInstructionExecution(ctx context.Context, processTestInstructionExecutionReveredRequest *fenixExecutionConnectorGrpcApi.ProcessTestInstructionExecutionReveredRequest) (finalTestInstructionExecutionResultMessage *fenixExecutionConnectorGrpcApi.FinalTestInstructionExecutionResultMessage, err error) {

	s.logger.WithFields(logrus.Fields{
		"id": "5a978baf-a4ab-402b-b36c-dc3615a8a6e9",
	}).Debug("Incoming 'gRPCServer - TriggerPostRestCallForTestInstructionExecution'")

	defer s.logger.WithFields(logrus.Fields{
		"id": "d822a6a1-8be5-4080-931b-5d9cf9771393",
	}).Debug("Outgoing 'gRPCServer - TriggerPostRestCallForTestInstructionExecution'")

	// Calling system
	userId := "External Trigger"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(userId, processTestInstructionExecutionReveredRequest.ProtoFileVersionUsedByClient)
	if returnMessage != nil {

		// Create TimeStamp in gRPC-format
		var grpcCurrentTimeStamp *timestamppb.Timestamp
		grpcCurrentTimeStamp = timestamppb.Now()

		finalTestInstructionExecutionResultMessage = &fenixExecutionConnectorGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: &fenixExecutionConnectorGrpcApi.ClientSystemIdentificationMessage{
				DomainUuid:                   string(Domains.DomainUUID_CA),
				ProtoFileVersionUsedByCaller: fenixExecutionConnectorGrpcApi.CurrentFenixExecutionConnectorProtoFileVersionEnum(common_config.GetHighestConnectorProtoFileVersion()),
			},
			TestInstructionExecutionUuid:         processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionUuid,
			TestInstructionExecutionStatus:       fenixExecutionConnectorGrpcApi.TestInstructionExecutionStatusEnum_TIE_CONTROLLED_INTERRUPTION,
			TestInstructionExecutionEndTimeStamp: grpcCurrentTimeStamp,
		}

		return finalTestInstructionExecutionResultMessage, nil
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
