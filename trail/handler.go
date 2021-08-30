package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	wg sync.WaitGroup
)

func panicHandler() {
	if err := recover(); err != nil {
		log.Println("Could not get predictions:", err)
	}
}

func writeOutput(src chan SLUResponse) {
	fp, err := os.OpenFile(outputCsv, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(fp)
	defer writer.Flush()
	defer fp.Close()

	for msg := range src {
		entities, err := json.Marshal(msg.Response.Entities)
		intents, err := json.Marshal(msg.Response.Intents)
		if err != nil {
			log.Fatalln("Could not Marshal ", err)
		}
		record := []string{msg.Uuid, string(entities), string(intents)}
		if err := writer.Write(record); err != nil {
			log.Fatalln("Error writing record to file", err)
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
	maxGoroutines := 10
	guard := make(chan struct{}, maxGoroutines)

	for _, item := range record.RequestBody {
		wg.Add(1)
		guard <- struct{}{} // would block if guard channel is already filled
		go func(outputChannel chan SLUResponse, item SLURequestBody) {
			go sluClient.Predict(outputChannel, item)
			<-guard
		}(outputChannel, item)
		// go sluClient.Predict(outputChannel, item)
	}

	go func() {
		wg.Wait()
		close(outputChannel)
	}()

	writeOutput(outputChannel)
	log.Printf("Predictions exported at: %s", outputCsv)
}
