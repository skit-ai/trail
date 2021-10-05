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

type CompleteResponse struct {
	Uuid     string                 `json:"uuid"`
	Response map[string]interface{} `json:"slu_response"`
}

type SLUResponse struct {
	Uuid     string `json:"uuid"`
	Response struct {
		Entities []struct {
			Text             string                 `json:"body"`
			Type             string                 `json:"type"`
			Score            float64                `json:"score"`
			Value            interface{}            `json:"value"`
			Values           []interface{}          `json:"values"`
			AlternativeIndex int                    `json:"alternative_index"`
			Dim              string                 `json:"dim"`
			EntityType       string                 `json:"entity_type"`
			Latent           bool                   `json:"latent"`
			Parsers          []string               `json:"parsers"`
			Range            map[string]interface{} `json:"range"`
			SlotNames        []string               `json:"slot_names"`
			Meta             struct{}               `json:"_meta"`
		} `json:"entities"`
		Intents []struct {
			Name             string        `json:"name"`
			AlternativeIndex int           `json:"alternative_index"`
			Parsers          []interface{} `json:"parsers"`
			Score            float64       `json:"score"`
			Slots            interface{}   `json:"slots"`
		} `json:"intents"`
		// Entities []interface{} `json:"entities"`
		// Intents  []interface{} `json:"intents"`
	} `json:"response"`
}

func NewSLUClient(host string) *SLU {
	return &SLU{
		HOST: host,
	}
}

// Call SLU service /predict/ endpoint
func (slu *SLU) Predict(outputChannel chan SLUResponse, sluRequestBody SLURequestBody) (SLUResponse, CompleteResponse) {
	defer panicHandler()
	jsonData, err := json.Marshal(sluRequestBody)
	method := "POST"
	client := &http.Client{}
	requestUrl := fmt.Sprintf(predict_api, slu.HOST, sluLanguage, sluClient)
	sluResponse := SLUResponse{
		Uuid: sluRequestBody.Context.Uuid,
	}
	completeResponse := CompleteResponse{
		Uuid: sluRequestBody.Context.Uuid,
	}

	req, err := http.NewRequest(method, requestUrl, bytes.NewReader(jsonData))
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return sluResponse, completeResponse
	}
	defer response.Body.Close()

	log.Printf("Getting predictions for UUID: %s ", sluRequestBody.Context.CallUuid)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseText map[string]interface{}
	err = json.Unmarshal(body, &sluResponse)
	err = json.Unmarshal(body, &responseText)

	completeResponse.Response = responseText

	if err != nil {
		log.Println("Could not unmarshal", err)
	}

	// outputChannel <- sluResponse
	// defer wg.Done()
	return sluResponse, completeResponse
}
