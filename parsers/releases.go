package parsers

import (
	"fmt"

	"github.com/google/go-github/v64/github"
	"github.com/hashicorp/go-version"
)

func ParseReleases(releases []*github.RepositoryRelease, afterVersion string) []string {
	currVersion, err := version.NewVersion(afterVersion)
	if err != nil {
		fmt.Println("Error creating new go-version: ", err.Error())
	}
	for _, rel := range releases {
		if *rel.Prerelease == false {
			relVersion, err := version.NewVersion(rel.GetTagName())
			if err != nil {
				fmt.Println("Error creating new go-version from release: ", err.Error())
			}
			if relVersion.GreaterThan(currVersion) {
				fmt.Println(rel.GetTagName())
			}
		}
	}
	return []string{}
}
