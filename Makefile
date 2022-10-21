RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6672

BuildExeForWindows:
#	fyne-cross windows -arch=amd64 --ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X 'main.gcpAuthentication=false'"
#	GOOD=windows GOARCH=amd64 go build -o FenixCAConnectorWindow.exe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X  'main.gcpAuthentication=true' -X 'main.caEngineAddress=127.0.0.1' -X 'main.caEngineAddressPath=/"
	env GOOD=windows GOARCH=amd64 go build -o FenixCAConnectorWindow.exe -ldflags="-X main.useInjectedEnvironmentVariables=true -X main.runInTray=truex -X main.loggingLevel=DebugLevel -X main.executionConnectorPort=6672 -X main.executionLocationForConnector=LOCALHOST_NODOCKER -X main.executionLocationForWorker=GCP -X main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app -X main.executionWorkerPort=443 -X  main.gcpAuthentication=true -X main.caEngineAddress=127.0.0.1 -X main.caEngineAddressPath=/"

BuildExeForLinux:
	GOOD=linux GOARCH=amd64 go build -ldflags="-X main.useInjectedEnvironmentVariales=true -X main.runInTray=true -X main.loggingLevel=DebugLevel -X main.executionConnectorPort=6672 -X main.executionLocationForConnector=LOCALHOST_NODOCKER -X main.executionLocationForWorker=LOCALHOST_NODOCKER -X main.executionWorkerAddress=127.0.0.1 -X main.executionWorkerPort=6671" -o FenixCAConnector.LinuxExe

