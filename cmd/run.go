package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zigzter/break-check/api"
	"github.com/zigzter/break-check/parsers"
)

var exclude bool

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the program",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.GetGHClient()
		packageJson, err := parsers.ParsePackageJSON()
		if err != nil {
			fmt.Println("Error parsing package.json: ", err.Error())
		}
		filterMap := make(map[string]bool)
		for _, dep := range args {
			filterMap[dep] = true
		}
		packages := parsers.ParsePackageVersions(packageJson, filterMap, exclude)
		repos := api.GetRepoNames(cmd.Context(), packages)
		for _, repo := range repos {
			releases, _, err := client.Repositories.ListReleases(cmd.Context(), repo.Owner, repo.Name, nil)
			if err != nil {
				fmt.Println("Error getting releases: ", err.Error())
			}
			parsers.ParseReleases(releases, "v16.13.1")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&exclude, "exclude", "e", false, "Exclude the provided packages from the check")
}
