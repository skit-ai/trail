package main

import (
	"github.com/spf13/cobra"
)

var (
	inputCsv    string
	outputCsv   string
	sluHost     string
	sluClient   string
	sluLanguage string
)

func main() {

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
	cmdFollow.PersistentFlags().StringVar(&inputCsv, "input-csv", "", "input csv file (required)")
	cmdFollow.MarkPersistentFlagRequired("input-csv")

	cmdFollow.PersistentFlags().StringVar(&outputCsv, "output-csv", "", "output csv file (required)")
	cmdFollow.MarkPersistentFlagRequired("output-csv")

	cmdFollow.PersistentFlags().StringVar(&sluHost, "slu-host", "", "http://host:port for SLU service (required)")
	cmdFollow.MarkPersistentFlagRequired("slu-host")

	cmdFollow.PersistentFlags().StringVar(&sluClient, "slu-client", "", "Name of the client (required)")
	cmdFollow.MarkPersistentFlagRequired("slu-client")

	cmdFollow.PersistentFlags().StringVar(&sluLanguage, "slu-language", "", "Language code. Example: en, hi (required)")
	cmdFollow.MarkPersistentFlagRequired("slu-language")

	rootCmd.AddCommand(cmdFollow)
	rootCmd.Execute()
}
