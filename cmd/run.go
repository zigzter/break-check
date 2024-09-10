package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the program",
	Run: func(cmd *cobra.Command, args []string) {
		exclude, err := cmd.Flags().GetBool("exclude")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Running, exclude set to: %t, args: %s\n", exclude, strings.Join(args, ","))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("exclude", "e", false, "Exclude the provided packages from the check")
}
