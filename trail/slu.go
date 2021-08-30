// All SLU related functions go here

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	predict_api = "%s/predict/%s/%s/"
)

type SLU struct {
	HOST string
}

type SLUResponse struct {
	Uuid     string `json:"uuid"`
	Response struct {
		Entities []interface{} `json:"entities"`
		Intents  []interface{} `json:"intents"`
	} `json:"response"`
}

func NewSLUClient(host string) *SLU {
	return &SLU{
		HOST: host,
	}
}

func (slu *SLU) Predict(outputChannel chan SLUResponse, sluRequestBody SLURequestBody) {
	jsonData, err := json.Marshal(sluRequestBody)
	requestUrl := fmt.Sprintf(predict_api, slu.HOST, sluLanguage, sluClient)
	response, err := http.Post(requestUrl, "application/json", bytes.NewReader(jsonData))

	log.Printf("Getting predictions for UUID: %s ", sluRequestBody.Context.CallUuid)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	sluResponse := SLUResponse{
		Uuid: sluRequestBody.Context.Uuid,
	}
	err = json.Unmarshal(body, &sluResponse)
	if err != nil {
		log.Fatal("error happened", err)
	}

	outputChannel <- sluResponse
	defer wg.Done()
}
