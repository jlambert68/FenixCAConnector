RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6672

BuildExeForWindows:
	fyne-cross windows -arch=amd64 -ldflags="-X main.useInjectedEnvironmentVariales=true -X main.runInTray=true -X main.loggingLevel=DebugLevel -X main.executionConnectorPort=6672 -X main.executionLocationForConnector=LOCALHOST_NODOCKER -X main.executionLocationForWorker=LOCALHOST_NODOCKER -X main.executionWorkerAddress=127.0.0.1 -X main.executionWorkerPort=6671"make

BuildExeForLinux:
	GOOD=linux GOARCH=amd64 go build -ldflags="-X main.useInjectedEnvironmentVariales=true -X main.runInTray=true -X main.loggingLevel=DebugLevel -X main.executionConnectorPort=6672 -X main.executionLocationForConnector=LOCALHOST_NODOCKER -X main.executionLocationForWorker=LOCALHOST_NODOCKER -X main.executionWorkerAddress=127.0.0.1 -X main.executionWorkerPort=6671" -o FenixCAConnector.LinuxExe

