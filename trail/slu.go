// All SLU related functions go here

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	services "***REMOVED***/services"
	dataframes "trail/trail/protos"
)

var (
    labledDataframe dataframes.LabeledDataFrame
)

func Predict(jsonMap map[string]interface{}) {
	plan, _ := ioutil.ReadFile("request.json")
    err := json.Unmarshal(plan, &jsonMap)
	if err != nil {
		log.Fatal("Unable to unmarshal json")
	}

	sluClient := services.Plute{HOST: "http://localhost:6969"}
	sluResponse, predictionResponse, _, err := sluClient.Predict("en", "indigo_slu", jsonMap, true)
	if err != nil {
		log.Fatal(err)
	}
    for idx, intent := range sluResponse.Prediction.Intents {
	    fmt.Println(idx, intent.Name)
    }
	fmt.Println(predictionResponse.Prediction.RawMessage)
}
