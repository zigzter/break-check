package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v64/github"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/spf13/cobra"
)

var ghClient *github.Client

func getGHClient() *github.Client {
	if ghClient == nil {
		// TODO: Set proper path for cache
		transport := httpcache.NewTransport(diskcache.New("."))
		transport.MarkCachedResponses = true
		client := &http.Client{
			Transport: transport,
		}
		ghClient = github.NewClient(client)
	}
	return ghClient
}

type repoInfo struct {
	Owner string
	Name  string
}

func getRepos(ctx context.Context, packages []string) []repoInfo {
	client := getGHClient()
	results := []repoInfo{}
	for _, pkg := range packages {
		// TODO: sanitize name, remove slash trailing subpackages
		result, res, err := client.Search.Repositories(ctx, pkg, nil)
		fmt.Println(res)
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
		repos := getRepos(cmd.Context(), args)
		for _, repo := range repos {
			fmt.Println(repo.Owner, repo.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("exclude", "e", false, "Exclude the provided packages from the check")
}
