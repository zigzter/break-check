package api

import (
	"context"
	"fmt"
)

type repoInfo struct {
	Owner string
	Name  string
}

func GetRepoNames(ctx context.Context, packages []string) []repoInfo {
	client := GetGHClient()
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
