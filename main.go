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

		case "UseInternalWebServerForTest":
			environmentVariable = useInternalWebServerForTest

		case "UseServiceAccount":
			environmentVariable = useServiceAccount

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
	useInternalWebServerForTest     string
	useServiceAccount               string
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

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
		_, returnMessageAckNack, returnMessageString := gcp.Gcp.GenerateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			// If there was any problem then exit program
			log.Fatalf(fmt.Sprintln("Couldn't generate access token for GCP, return message: '%s'", returnMessageString))
		}
	}

	// InitiateRestCallsToCAEngine()
	restCallsToCAEngine.InitiateRestCallsToCAEngine()

	// If local web server, used for testing, should be used instead of FangEngine
	if common_config.UseInternalWebServerForTest == true {

		common_config.Logger.WithFields(logrus.Fields{
			"id": "353930b1-5c6f-4826-955c-19f543e2ab85",
		}).Info("Using internal web server instead of FangEngine, for RestCall")

		go func() {

			restCallsToCAEngine.RestAPIServer()
			/*

				type jsonType string

				http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					if r.Method != http.MethodPost {
						http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
						return
					}

					fmt.Println("got r.Body:", r.Body)

					// read response body
					body, error := ioutil.ReadAll(r.Body)
					if error != nil {
						fmt.Println(error)
					}
					// close response body
					r.Body.Close()

					var myjson jsonType

					jsonMap := make(map[string]interface{})
					err := json.Unmarshal(body, &jsonMap)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					fmt.Println(jsonMap)

					fmt.Println("got r.Body:", r.Body)
					fmt.Println("myjson:", myjson)
					w.WriteHeader(http.StatusOK)


			*/

			// ***************************************
			/*
						err := json.NewDecoder(r.Body).Decode(myjson)
						if err != nil {
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}



					/*
							jsonMap := make(map[string]interface{})
							err := json.Unmarshal([]byte(jsonStr), &jsonMap)
							if err != nil {
								panic(err)
							}
							dumpMap("", jsonMap)



						var data map[string]interface{}
						err := json.Unmarshal([]byte(r.body), &data)
						if err != nil {
							fmt.Println("Couldn't Unmarshal Rest-body")
						}

					fmt.Println("got r.Body:", r.Body)
					fmt.Println("myjson:", myjson)
					w.WriteHeader(http.StatusOK)


				})

			*/

			//if err := http.ListenAndServe(common_config.LocalWebServerAddressAndPort, nil); err != http.ErrServerClosed {
			//	panic(err)
			//}
		}()
	}

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
					newNotification := fyne.NewNotification("Fenix Execution Connector", "Fenix will rule the 'Test World'")

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

				time.Sleep(time.Second * 7)
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

	// Extract if local web server for test should be used instead of FangEngine
	boolValue, err = strconv.ParseBool(mustGetenv("UseInternalWebServerForTest"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseInternalWebServerForTest:' to an boolean, error: ", err)
		os.Exit(0)
	}
	common_config.UseInternalWebServerForTest = boolValue

	// Extract Address to Custody Arrangement Rest-Engine
	common_config.CAEngineAddress = mustGetenv("CAEngineAddress")
	common_config.CAEngineAddressPath = mustGetenv("CAEngineAddressPath")

	// Extract if Service Account should be used towards GCP or should the user log in via web
	boolValue, err = strconv.ParseBool(mustGetenv("UseServiceAccount"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseServiceAccount:' to an boolean, error: ", err)
		os.Exit(0)
	}
	common_config.UseServiceAccount = boolValue

}
