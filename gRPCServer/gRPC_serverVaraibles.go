package gRPCServer

import (
	"FenixExecutionWorker/connectorEngine"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type FenixExecutionConnectorGrpcObjectStruct struct {
	logger                       *logrus.Logger
	ExecutionConnectorGrpcObject *FenixExecutionConnectorGrpcObjectStruct
}

// Variable holding everything together
//var ExecutionConnectorGrpcObject *FenixExecutionConnectorGrpcObjectStruct

// gRPCServer variables
var (
	fenixExecutionConnectorGrpcServer                       *grpc.Server
	registerFenixExecutionConnectorGrpcServicesServer       *grpc.Server
	registerFenixExecutionConnectorWorkerGrpcServicesServer *grpc.Server
	lis                                                     net.Listener
)

// gRPCServer Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type fenixExecutionConnectorGrpcServicesServer struct {
	logger                  *logrus.Logger
	CommandChannelReference *connectorEngine.ExecutionEngineChannelType
	fenixExecutionWorkerGrpcApi.UnimplementedFenixExecutionWorkerGrpcServicesServer
}

// gRPCServer Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type fenixExecutionConnectorWorkerGrpcServicesServer struct {
	logger                  *logrus.Logger
	CommandChannelReference *connectorEngine.ExecutionEngineChannelType
	fenixExecutionWorkerGrpcApi.UnimplementedFenixExecutionWorkerConnectorGrpcServicesServer
}
