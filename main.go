package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bc",
	Short: "Break Check collects all breaking changes in the given packages up to the latest version",
	Long:  "Break Check collects all breaking changes in the given packages up to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("break-check")
	},
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
