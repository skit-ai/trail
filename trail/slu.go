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

// Call SLU service /predic/ endpoint
func (slu *SLU) Predict(outputChannel chan SLUResponse, sluRequestBody SLURequestBody) SLUResponse {
	defer panicHandler()
	jsonData, err := json.Marshal(sluRequestBody)
	method := "POST"
	client := &http.Client{}
	requestUrl := fmt.Sprintf(predict_api, slu.HOST, sluLanguage, sluClient)
	sluResponse := SLUResponse{
		Uuid: sluRequestBody.Context.Uuid,
	}

	req, err := http.NewRequest(method, requestUrl, bytes.NewReader(jsonData))
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return sluResponse
	}
	defer response.Body.Close()

	log.Printf("Getting predictions for UUID: %s ", sluRequestBody.Context.CallUuid)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &sluResponse)
	if err != nil {
		log.Println("Could not unmarshal", err)
	}

	// outputChannel <- sluResponse
	// defer wg.Done()
	return sluResponse
}
