package restCallsToCAEngine

import (
	"FenixCAConnector/common_config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func postExecutionInstruction(todo Todo) (err error) {
	fmt.Println("2. Performing Http Post...")

	common_config.Logger.WithFields(logrus.Fields{
		"id":   "38c5fd40-0aee-4cd0-9107-0974331db0cc",
		"todo": todo,
	}).Debug("Posting ExecutionInstruction to Custody Arrangements execution Engine")

	todo = Todo{1, 2, "lorem ipsum dolor sit amet", true}

	jsonReq, err := json.Marshal(todo)
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"id":   "e1e74131-5040-43fa-abfc-1023f09d4388",
			"todo": todo,
		}).Error("Couldn't Marshal into json request")

		return err
	}

	resp, err := http.Post(common_config.CAEngineAddress+common_config.CAEngineAddressPath, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"id": "b98c2fb4-e717-4fc4-8d2c-6c791c523175",
			"common_config.CAEngineAddress + common_config.CAEngineAddressPath": common_config.CAEngineAddress + common_config.CAEngineAddressPath,
		}).Error("Couldn't do call to Custody Arrangements Rest-execution-server")

		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)

	return err
}
