package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v64/github"
	"github.com/spf13/cobra"
)

var ghClient *github.Client

func getGHClient() *github.Client {
	if ghClient == nil {
		ghClient = github.NewClient(nil)
	}
	return ghClient
}

type repoInfo struct {
	Owner string
	Name  string
}

func getRepos(packages []string) []repoInfo {
	client := getGHClient()
	results := make([]repoInfo, len(packages))
	for _, pkg := range packages {
		// TODO: sanitize name, remove slash trailing subpackages
		result, _, err := client.Search.Repositories(context.TODO(), pkg, nil)
		if err != nil {
			fmt.Println("search error: ", err.Error())
			continue
		}
		if len(result.Repositories) > 0 {
			firstResult := result.Repositories[0]
			results = append(results, repoInfo{Name: *firstResult.Name, Owner: *firstResult.Owner.Login})
		}
	}
	return results
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the program",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepos(args)
		for _, repo := range repos {
			fmt.Println(repo.Owner, repo.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("exclude", "e", false, "Exclude the provided packages from the check")
}
