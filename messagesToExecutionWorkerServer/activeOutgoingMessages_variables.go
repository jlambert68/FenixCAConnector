package messagesToExecutionWorkerServer

import (
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

type MessagesToExecutionWorkerObjectStruct struct {
	logger         *logrus.Logger
	gcpAccessToken *oauth2.Token
	//CommandChannelReference *connectorEngine.ExecutionEngineChannelType
}

// Variables used for contacting Fenix Execution Worker Server
var (
	remoteFenixExecutionWorkerServerConnection *grpc.ClientConn
	FenixExecutionWorkerAddressToDial          string
	fenixExecutionWorkerGrpcClient             fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServicesClient
)
