package messagesToExecutionWorkerServer

import (
	"FenixCAConnector/common_config"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ********************************************************************************************************************

// SetConnectionToFenixExecutionWorkerServer - Set upp connection and Dial to FenixExecutionServer
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) SetConnectionToFenixExecutionWorkerServer() (err error) {

	var opts []grpc.DialOption

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		panic(fmt.Sprintf("cannot load root CA certs, err: %s", err))
	}

	//When running on GCP then use credential otherwise not
	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            systemRoots,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to Fenix Execution Worker
	// When run on GCP, use credentials
	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP {
		// Run on GCP
		remoteFenixExecutionWorkerServerConnection, err = grpc.Dial(common_config.FenixExecutionWorkerAddressToDial, opts...)
	} else {
		// Run Local
		remoteFenixExecutionWorkerServerConnection, err = grpc.Dial(common_config.FenixExecutionWorkerAddressToDial, grpc.WithInsecure())
	}
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"ID": "50b59b1b-57ce-4c27-aa84-617f0cde3100",
			"common_config.FenixExecutionWorkerAddressToDial": common_config.FenixExecutionWorkerAddressToDial,
			"error message": err,
		}).Error("Did not connect to FenixExecutionServer via gRPC")

		return err

	} else {
		common_config.Logger.WithFields(logrus.Fields{
			"ID": "0c650bbc-45d0-4029-bd25-4ced9925a059",
			"common_config.FenixExecutionWorkerAddressToDial": common_config.FenixExecutionWorkerAddressToDial,
		}).Info("gRPC connection OK to FenixExecutionServer")

		// Creates a new Client
		fenixExecutionWorkerGrpcClient = fenixExecutionWorkerGrpcApi.NewFenixExecutionWorkerConnectorGrpcServicesClient(remoteFenixExecutionWorkerServerConnection)

	}
	return err
}

/*
// Generate Google access token. Used when running in GCP
func (toExecutionWorkerObject *MessagesToExecutionWorkerObjectStruct) generateGCPAccessToken(ctx context.Context) (appendedCtx context.Context, returnAckNack bool, returnMessage string) {

	// Only create the token if there is none, or it has expired
	if toExecutionWorkerObject.GcpAccessToken == nil || toExecutionWorkerObject.GcpAccessToken.Expiry.Before(time.Now()) {

		// Create an identity token.
		// With a global TokenSource tokens would be reused and auto-refreshed at need.
		// A given TokenSource is specific to the audience.
		tokenSource, err := idtoken.NewTokenSource(ctx, "https://"+common_config.FenixExecutionWorkerAddress)
		if err != nil {
			common_config.Logger.WithFields(logrus.Fields{
				"ID":  "8ba622d8-b4cd-46c7-9f81-d9ade2568eca",
				"err": err,
			}).Error("Couldn't generate access token")

			return nil, false, "Couldn't generate access token"
		}

		token, err := tokenSource.Token()
		if err != nil {
			common_config.Logger.WithFields(logrus.Fields{
				"ID":  "0cf31da5-9e6b-41bc-96f1-6b78fb446194",
				"err": err,
			}).Error("Problem getting the token")

			return nil, false, "Problem getting the token"
		} else {
			common_config.Logger.WithFields(logrus.Fields{
				"ID":    "8b1ca089-0797-4ee6-bf9d-f9b06f606ae9",
				"token": token,
			}).Debug("Got Bearer Token")
		}

		toExecutionWorkerObject.GcpAccessToken = token

	}

	common_config.Logger.WithFields(logrus.Fields{
		"ID": "cd124ca3-87bb-431b-9e7f-e044c52b4960",
		"FenixExecutionWorkerObject.gcpAccessToken": toExecutionWorkerObject.GcpAccessToken,
	}).Debug("Will use Bearer Token")

	// Add token to GrpcServer Request.
	appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+toExecutionWorkerObject.GcpAccessToken.AccessToken)

	return appendedCtx, true, ""

}

*/
