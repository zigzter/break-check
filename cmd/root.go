package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bc",
	Short: "Break Check collects all breaking changes in the given packages up to the latest version",
	Long:  "Break Check collects all breaking changes in the given packages up to the latest version",
}

func Execute() error {
	return rootCmd.Execute()
}
