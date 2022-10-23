package common_config

import "github.com/sirupsen/logrus"

// Used for keeping track of the proto file versions for ExecutionServer and this Worker
var highestConnectorProtoFileVersion int32 = -1
var highestExecutionWorkerProtoFileVersion int32 = -1

// Logger that all part of the system can use
var Logger *logrus.Logger

const LocalWebServerAddressAndPort = ":8080"
