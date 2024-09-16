package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zigzter/break-check/api"
	"github.com/zigzter/break-check/parsers"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the program",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.GetGHClient()
		repos := api.GetRepoNames(cmd.Context(), args)
		for _, repo := range repos {
			releases, _, err := client.Repositories.ListReleases(cmd.Context(), repo.Owner, repo.Name, nil)
			if err != nil {
				fmt.Println("Error getting releases: ", err.Error())
			}
			changes := parsers.ParseRelease(*releases[4])
			for _, change := range changes {
				fmt.Println(change)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("exclude", "e", false, "Exclude the provided packages from the check")
}
