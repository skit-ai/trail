package main

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

// rootHandler - follow command handler
func rootHandler(cmd *cobra.Command, args []string) {
	reader := Reader{path: inputProto}
	record := reader.yamlReader()

    // log record to stdout
	for idx, element := range record.Calls {
		fmt.Println(idx, element.Id)
	}

	jsonMap := make(map[string]interface{})

    var wg sync.WaitGroup
    wg.Add(1)
    go Predict(jsonMap)
    wg.Wait()
}
