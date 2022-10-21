package main

import (
	"FenixCAConnector/common_config"
	"FenixCAConnector/gcp"
	"FenixCAConnector/restCallsToCAEngine"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	//"flag"
	"fmt"
	"log"
	"os"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(environmentVariableName string) string {

	var environmentVariable string

	if useInjectedEnvironmentVariables == "true" {
		// Extract environment variables from parameters feed into program at compilation time

		switch environmentVariableName {
		case "RunInTray":
			environmentVariable = runInTray
		case "LoggingLevel":
			environmentVariable = loggingLevel

		case "ExecutionConnectorPort":
			environmentVariable = executionConnectorPort

		case "ExecutionLocationForConnector":
			environmentVariable = executionLocationForConnector

		case "ExecutionLocationForWorker":
			environmentVariable = executionLocationForWorker

		case "ExecutionWorkerAddress":
			environmentVariable = executionWorkerAddress

		case "ExecutionWorkerPort":
			environmentVariable = executionWorkerPort

		case "GCPAuthentication":
			environmentVariable = gcpAuthentication

		case "CAEngineAddress":
			environmentVariable = caEngineAddress

		case "CAEngineAddressPath":
			environmentVariable = caEngineAddressPath

		default:
			log.Fatalf("Warning: %s environment variable not among injected variables.\n", environmentVariableName)

		}

		if environmentVariable == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", environmentVariableName)
		}

	} else {
		//
		environmentVariable = os.Getenv(environmentVariableName)
		if environmentVariable == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", environmentVariableName)
		}

	}
	return environmentVariable
}

// Variables injected at compilation time
var (
	useInjectedEnvironmentVariables string
	runInTray                       string
	loggingLevel                    string
	executionConnectorPort          string
	executionLocationForConnector   string
	executionLocationForWorker      string
	executionWorkerAddress          string
	executionWorkerPort             string
	gcpAuthentication               string
	caEngineAddress                 string
	caEngineAddressPath             string
)

func main() {

	var logFileName string

	// Extract from environment variables if it should run as a tray application or not
	shouldRunInTray := mustGetenv("RunInTray")

	// When run as Tray application then add log-name
	if shouldRunInTray == "true" {
		logFileName = "fenixConnectorLog.log"
	} else {
		logFileName = ""
	}

	// Initiate logger in common_config
	InitLogger(logFileName)

	// When Execution Worker runs on GCP, then set up access
	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP && common_config.GCPAuthentication == true {
		gcp.Gcp = gcp.GcpObjectStruct{}

		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// Generate first time Access token
		_, returnMessageAckNack, returnMessageString := gcp.Gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			// If there was any problem then exit program
			log.Fatalf(fmt.Sprintln("Couldn't generate access token for GCP, return message: '%s'", returnMessageString))
		}
	}

	// InitiateRestCallsToCAEngine()
	restCallsToCAEngine.InitiateRestCallsToCAEngine()

	// Start Connector Engine
	go fenixExecutionConnectorMain()

	// Start up Tray Application if it should that as that
	if shouldRunInTray == "true" {
		// Start application as TrayApplication

		a := app.NewWithID("FenixCAConnector")
		a.SetIcon(resourceFenix57Png)
		mainFyneWindow := a.NewWindow("SysTray")

		if desk, ok := a.(desktop.App); ok {
			m := fyne.NewMenu("Fenix Execution Connector",
				fyne.NewMenuItem("Hide", func() {
					mainFyneWindow.Hide()
					newNotification := fyne.NewNotification("MyTitle", "MyCOntent")

					a.SendNotification(newNotification)
				}))
			desk.SetSystemTrayMenu(m)
		}

		// Create Fenix Splash screen
		var splashWindow fyne.Window
		if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
			splashWindow = drv.CreateSplashWindow()

			// Fenix Header
			fenixHeaderText := canvas.Text{
				Alignment: fyne.TextAlignCenter,
				Color:     nil,
				Text:      "Fenix Inception - SaaS",
				TextSize:  20,
				TextStyle: fyne.TextStyle{Bold: true},
			}

			// Text Footer
			halFinney := widget.NewLabel("\"If you want to change the world, don't protest. Write code!\" - Hal Finney (1994)")

			// Fenix picture
			image := canvas.NewImageFromResource(resourceFenix12Png)
			image.FillMode = canvas.ImageFillOriginal

			// Container holding Header, picture and Footer
			spashContainer := container.New(layout.NewVBoxLayout(), &fenixHeaderText, image, halFinney)

			splashWindow.SetContent(spashContainer)
			splashWindow.CenterOnScreen()
			splashWindow.Show()

			go func() {
				time.Sleep(time.Millisecond * 1000)

				mainFyneWindow.Hide()

				time.Sleep(time.Second * 6)
				splashWindow.Close()

			}()

			mainFyneWindow.SetContent(widget.NewLabel("Fyne System Tray"))
			mainFyneWindow.SetCloseIntercept(func() {
				mainFyneWindow.Hide()
			})

			mainFyneWindow.Hide()
			go func() {
				count := 10
				for {
					time.Sleep(time.Millisecond * 100)
					//mainFyneWindow.Hide()
					count = count - 1
					if count == 0 {
						break
					}
					return
				}

				//mainFyneWindow.Hide()
			}()

			mainFyneWindow.ShowAndRun()
		}

	} else {
		// Run as console program and exit as on standard exiting signals
		sig := make(chan os.Signal, 1)
		done := make(chan bool, 1)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sig
			fmt.Println()
			fmt.Println(sig)
			done <- true

			fmt.Println("ctrl+c")
		}()

		fmt.Println("awaiting signal")
		<-done
		fmt.Println("exiting")
	}
}

func init() {
	//executionLocationForConnector := flag.String("startupType", "0", "The application should be started with one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
	//flag.Parse()

	var err error

	// Get Environment variable to tell how/were this worker is  running
	var executionLocationForConnector = mustGetenv("ExecutionLocationForConnector")

	switch executionLocationForConnector {
	case "LOCALHOST_NODOCKER":
		common_config.ExecutionLocationForConnector = common_config.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		common_config.ExecutionLocationForConnector = common_config.LocalhostDocker

	case "GCP":
		common_config.ExecutionLocationForConnector = common_config.GCP

	default:
		fmt.Println("Unknown Execution location for Connector: " + executionLocationForConnector + ". Expected one of the following: 'LOCALHOST_NODOCKER', 'LOCALHOST_DOCKER', 'GCP'")
		os.Exit(0)

	}

	// Get Environment variable to tell were Fenix Execution Server is running
	var executionLocationForExecutionWorker = mustGetenv("ExecutionLocationForWorker")

	switch executionLocationForExecutionWorker {
	case "LOCALHOST_NODOCKER":
		common_config.ExecutionLocationForFenixExecutionWorkerServer = common_config.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		common_config.ExecutionLocationForFenixExecutionWorkerServer = common_config.LocalhostDocker

	case "GCP":
		common_config.ExecutionLocationForFenixExecutionWorkerServer = common_config.GCP

	default:
		fmt.Println("Unknown Execution location for Fenix Execution Worker Server: " + executionLocationForExecutionWorker + ". Expected one of the following: 'LOCALHOST_NODOCKER', 'LOCALHOST_DOCKER', 'GCP'")
		os.Exit(0)

	}

	// Address to Fenix Execution Worker Server
	common_config.FenixExecutionWorkerAddress = mustGetenv("ExecutionWorkerAddress")

	// Port for Fenix Execution Worker Server
	common_config.FenixExecutionWorkerPort, err = strconv.Atoi(mustGetenv("ExecutionWorkerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'ExecutionWorkerPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Port for Fenix Execution Connector Server
	common_config.ExecutionConnectorPort, err = strconv.Atoi(mustGetenv("ExecutionConnectorPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'executionConnectorPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Build the Dial-address for gPRC-call
	common_config.FenixExecutionWorkerAddressToDial = common_config.FenixExecutionWorkerAddress + ":" + strconv.Itoa(common_config.FenixExecutionWorkerPort)

	// Extract Debug level
	var loggingLevel = mustGetenv("LoggingLevel")

	switch loggingLevel {

	case "DebugLevel":
		common_config.LoggingLevel = logrus.DebugLevel

	case "InfoLevel":
		common_config.LoggingLevel = logrus.InfoLevel

	default:
		fmt.Println("Unknown loggingLevel '" + loggingLevel + "'. Expected one of the following: 'DebugLevel', 'InfoLevel'")
		os.Exit(0)

	}

	// Extract if there is a need for authentication when going toward GCP
	boolValue, err := strconv.ParseBool(mustGetenv("GCPAuthentication"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'GCPAuthentication:' to an boolean, error: ", err)
		os.Exit(0)
	}
	common_config.GCPAuthentication = boolValue

	// Extract Address to Custody Arrangement Rest-Engine
	common_config.CAEngineAddress = mustGetenv("CAEngineAddress")
	common_config.CAEngineAddressPath = mustGetenv("CAEngineAddressPath")

}
