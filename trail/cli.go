package main

import (
	"github.com/spf13/cobra"
)

var (
	inputProto string
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
	cmdFollow.PersistentFlags().StringVar(&inputProto, "input-proto", "", "input protobuf file")
	rootCmd.AddCommand(cmdFollow)
	rootCmd.Execute()
}
