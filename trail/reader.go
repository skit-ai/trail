package main

import (
	"encoding/csv"
	"encoding/json"
	"io"

	// "encoding/json"

	"io/ioutil"
	"log"
	"os"
	pb "trail/trail/protos"

	"google.golang.org/protobuf/proto"
	// fsm "***REMOVED***/services"
)

type Reader struct {
	path string
}

type SLURequestBody struct {
	Text         string        `json:"text"`
	Alternatives []interface{} `json:"alternatives"`
	Context      struct {
		CallUuid     string `json:"call_uuid"`
		CurrentState string `json:"current_state"`
		Uuid         string `json:"uuid"`
	} `json:"context"`
}

type ListRequests struct {
	RequestBody []SLURequestBody
}

func (reader Reader) yamlReader() *pb.CallDataFrame {
	in, err := ioutil.ReadFile(reader.path)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	record := &pb.CallDataFrame{}
	if err := proto.Unmarshal(in, record); err != nil {
		log.Fatalln("Failed to parse call dataframe:", err)
	}

	return record
}

func (reader Reader) csvReader() ListRequests {
	f, err := os.Open(reader.path)
	if err != nil {
		log.Fatal("Unable to read input file "+reader.path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	requests := ListRequests{}
	rowIdx := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if rowIdx == 0 {
			rowIdx++
			continue
		}

		sluRequest := SLURequestBody{}

		// Get alternatives
		var data []interface{}
		err = json.Unmarshal([]byte(record[2]), &data)
		if err != nil {
			log.Fatal("error happened", err)
		}
		sluRequest.Alternatives = data

		// Context
		sluRequest.Context.CallUuid = record[0]
		sluRequest.Context.CurrentState = record[6]
		sluRequest.Context.Uuid = record[1]

		requests.RequestBody = append(requests.RequestBody, sluRequest)
	}
	return requests
}
