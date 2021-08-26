package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "trail",
		Short: "Trail a path for a dataframe",
		Long:  `Tool to follow a dataframe and get predicted dataframes against your service config`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
