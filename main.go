package main

import (
	"FenixCAConnector/common_config"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
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
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

var runInTray string

func main() {

	var logFileName string
	// When run as Tray application then add log-name
	if runInTray == "true" {
		logFileName = "fenixConnectorLog.log"
	} else {
		logFileName = ""
	}

	go fenixExecutionConnectorMain(logFileName)

	if runInTray == "true" {
		// Start application as TrayApplication

		a := app.New()
		a.SetIcon(resourceFenix57Png)
		mainFyneWindow := a.NewWindow("SysTray")

		if desk, ok := a.(desktop.App); ok {
			m := fyne.NewMenu("Fenix Execution Connector",
				fyne.NewMenuItem("Hide", func() {
					mainFyneWindow.Hide()
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
				time.Sleep(time.Second * 7)
				splashWindow.Close()

			}()

			mainFyneWindow.SetContent(widget.NewLabel("Fyne System Tray"))
			mainFyneWindow.SetCloseIntercept(func() {
				mainFyneWindow.Hide()
			})

			go func() {
				time.Sleep(time.Millisecond * 1000)
				mainFyneWindow.Hide()
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
		fmt.Println("Couldn't convert environment variable 'ExecutionConnectorPort' to an integer, error: ", err)
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
		fmt.Println("Unknown LoggingLevel '" + loggingLevel + "'. Expected one of the following: 'DebugLevel', 'InfoLevel'")
		os.Exit(0)

	}

}
