package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	wg sync.WaitGroup
)

type StdOutResponse struct {
	Response []*SLUResponse
}

type CsvIntent struct {
	Name string
}

type CsvEntity struct {
	Text  string      `json:"text"`
	Type  string      `json:"type"`
	Score float64     `json:"score"`
	Value interface{} `json:"value"`
}

// Generic panic handler
func panicHandler() {
	if err := recover(); err != nil {
		log.Println("Could not get predictions:", err)
	}
}

// writeOutput - Write output from SLU service to CSV
func writeOutput(sluResponse []SLUResponse) {
	var intentsWriter, entitiesWriter *csv.Writer

	if outputIntentsCsv != "" {
		intentsFp, err := os.OpenFile(outputIntentsCsv, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer intentsFp.Close()
		intentsWriter = csv.NewWriter(intentsFp)
		defer intentsWriter.Flush()
		if err != nil {
			panic(err)
		}
	}
	if outputEntitiesCsv != "" {
		entitiesFp, err := os.OpenFile(outputEntitiesCsv, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer entitiesFp.Close()
		entitiesWriter = csv.NewWriter(entitiesFp)
		defer entitiesWriter.Flush()
		if err != nil {
			panic(err)
		}
	}

	for _, msg := range sluResponse {
		// Log intent in CSV
		if outputIntentsCsv != "" && len(msg.Response.Intents) >= 1 {
			// Get intent with max score
			var maxScore float64 = 0
			intentIdx := 0
			for idx, intent := range msg.Response.Intents {
				if intent.Score > maxScore {
					intentIdx = idx
                    maxScore = intent.Score
				}
			}
			csvIntent := CsvIntent{Name: msg.Response.Intents[intentIdx].Name}

			intentsRecord := []string{msg.Uuid, csvIntent.Name}
			if err := intentsWriter.Write(intentsRecord); err != nil {
				log.Fatalln("Error writing record to file", err)
			}
		}

		if outputEntitiesCsv != "" && len(msg.Response.Entities) >= 1 {
			csvEntityList := make([]CsvEntity, len(msg.Response.Entities))
			for idx, entity := range msg.Response.Entities {
				// value, _ := json.Marshal(entity.Value)
				csvEntity := CsvEntity{Text: entity.Text, Type: entity.Type, Score: entity.Score, Value: entity.Value}
				csvEntityList[idx] = csvEntity
			}

			entities, err := json.Marshal(csvEntityList)
			if err != nil {
				log.Fatalln("Could not Marshal ", err)
			}

			entitiesRecord := []string{msg.Uuid, string(entities)}
			if err := entitiesWriter.Write(entitiesRecord); err != nil {
				log.Fatalln("Error writing record to file", err)
			}
		}
	}
}

// rootHandler - follow command handler
func rootHandler(cmd *cobra.Command, args []string) {
	defer panicHandler()
	reader := Reader{path: inputCsv}
	record := reader.csvReader()
	sluClient := NewSLUClient(sluHost)
	outputChannel := make(chan SLUResponse)

	log.Println("Making requests to SLU service")
	stdOutResponse := make([]SLUResponse, len(record.RequestBody))

	guard := make(chan struct{}, maxGoroutines)

	for idx, item := range record.RequestBody {
		wg.Add(1)
		guard <- struct{}{} // would block if guard channel is already filled
		go func(outputChannel chan SLUResponse, item SLURequestBody, idx int) {
			sluResponse := sluClient.Predict(outputChannel, item)
			stdOutResponse[idx] = sluResponse

			defer wg.Done()
			<-guard
		}(outputChannel, item, idx)
	}

	wg.Wait()
	writeOutput(stdOutResponse)

	jsonData, _ := json.MarshalIndent(stdOutResponse, "", "    ")
	fmt.Println(string(jsonData))
}
