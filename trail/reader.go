package main

import (
	"encoding/csv"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
	pb "trail/trail/protos"
	// fsm "fsm-service/v3/services"
)

type Reader struct {
	path string
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

func (reader Reader) csvReader() [][]string {
	f, err := os.Open(reader.path)
	if err != nil {
		log.Fatal("Unable to read input file "+reader.path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+reader.path, err)
	}

	return records
}
