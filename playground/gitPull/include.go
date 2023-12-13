package gitPull

import (
	"github.com/mikebd/go-util/pkg/git"
	"slices"
)

var branchesOfInterest = []string{"main", "master"}

// includeGitRepositoryDirectoryCheap returns true if the specified directory
// should be included in the list of directories to pull.  A directory
// is included if its branch is one of the branchesOfInterest.
// This function includes only cheap tests that do not require goroutines.
func includeGitRepositoryDirectoryCheap(dir string) bool {
	globalOptions := git.GlobalOptions{AsIfIn: dir}
	branch, errBranch := git.CurrentBranchName(globalOptions)
	if errBranch != nil {
		return false
	}

	if !slices.Contains(branchesOfInterest, branch) {
		return false
	}

	return true
}
