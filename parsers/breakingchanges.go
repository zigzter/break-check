package parsers

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v64/github"
)

func isBreakingChangeHeader(str string) bool {
	lowercased := strings.ToLower(str)
	return strings.Contains(lowercased, "breaking") && strings.Contains(lowercased, "#")
}

func ParseBreakingChanges(release github.RepositoryRelease) []string {
	lines := strings.Split(*release.Body, "\n")
	isInBreakingChanges := false
	var breakingChangesList []string
	for _, curr := range lines {
		if isInBreakingChanges && strings.Contains(curr, "#") {
			fmt.Println("end of breaking changes")
			isInBreakingChanges = false
		} else if isInBreakingChanges {
			fmt.Println("within breaking changes")
			breakingChangesList = append(breakingChangesList, curr)
		} else if isBreakingChangeHeader(curr) {
			fmt.Println("found breaking changes header")
			isInBreakingChanges = true
		}
	}
	return breakingChangesList
}
