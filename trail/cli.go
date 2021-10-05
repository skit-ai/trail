package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	inputCsv          string
	outputIntentsCsv  string
	outputEntitiesCsv string
	sluHost           string
	sluClient         string
	sluLanguage       string
	recordType        string
	maxGoroutines     int
)

func main() {
	// version
	var version = &cobra.Command{
		Use:   "version",
		Short: "Show trail version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("0.3.0")
		},
	}

	// follow command
	var cmdFollow = &cobra.Command{
		Use:   "follow <args>",
		Short: "Follow a path for a dataframe",
		Long:  `Follow a dataframe and get predicted dataframes against your service config`,
		Run: func(cmd *cobra.Command, args []string) {
			rootHandler(cmd, args)
		},
	}
	var rootCmd = &cobra.Command{Use: "trail"}

	// Mandatory flags
	cmdFollow.PersistentFlags().StringVar(&inputCsv, "input-csv", "", "input csv file (required)")
	cmdFollow.MarkPersistentFlagRequired("input-csv")

	cmdFollow.PersistentFlags().StringVar(&sluHost, "slu-host", "", "http://host:port for SLU service (required)")
	cmdFollow.MarkPersistentFlagRequired("slu-host")

	cmdFollow.PersistentFlags().StringVar(&sluClient, "slu-client", "", "Name of the client (required)")
	cmdFollow.MarkPersistentFlagRequired("slu-client")

	cmdFollow.PersistentFlags().StringVar(&sluLanguage, "slu-language", "", "Language code. Example: en, hi (required)")
	cmdFollow.MarkPersistentFlagRequired("slu-language")

	// Optional flags
	cmdFollow.PersistentFlags().StringVar(&outputIntentsCsv, "output-intents-csv", "", "output intents csv file")
	cmdFollow.PersistentFlags().StringVar(&outputEntitiesCsv, "output-entities-csv", "", "output entities csv file")
	cmdFollow.PersistentFlags().IntVar(&maxGoroutines, "concurrency", 30, "Max concurrent requests to SLU service (optional)")
	cmdFollow.PersistentFlags().StringVar(&recordType, "type", "untagged", "Type of record. One of: [tagged, untagged] (optional)")

	rootCmd.AddCommand(cmdFollow)
	rootCmd.AddCommand(version)
	rootCmd.Execute()
}
