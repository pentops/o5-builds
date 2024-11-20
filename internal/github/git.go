package github

import (
	"regexp"
	"strings"

	"github.com/pentops/j5/gen/j5/source/v1/source_j5pb"
)

func ExpandGitAliases(latestBranch string, commitInfo *source_j5pb.CommitInfo) {
	aliases := make([]string, 0, len(commitInfo.Aliases))
	for _, alias := range commitInfo.Aliases {
		if strings.HasPrefix(alias, "refs/tags/") {
			aliases = append(aliases, strings.TrimPrefix(alias, "refs/tags/"))
		} else if strings.HasPrefix(alias, "refs/heads/") {
			branchName := strings.TrimPrefix(alias, "refs/heads/")
			aliases = append(aliases, branchName)
		} else {
			aliases = append(aliases, alias)
		}
		if globMatch(latestBranch, alias) {
			aliases = append(aliases, "latest")
		}
	}
	commitInfo.Aliases = aliases
}

func globMatch(pattern, s string) bool {
	escaped := regexp.QuoteMeta(pattern)
	// Replace escaped * with .* to make it a regexp pattern.
	pattern = strings.ReplaceAll(escaped, "\\*", ".*")
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return matcher.MatchString(s)
}
