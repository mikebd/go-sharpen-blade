package gitPull

import (
	"github.com/mikebd/go-util/pkg/git"
	"os"
	"slices"
)

var branchesOfInterest = []string{"main", "master"}

// remote is the repository short name that we want to track.
// This might become configurable in the future.
const remote = "origin"

func includeGitRepositoryDirectory(directory string) bool {
	currentDir, err := os.Getwd()
	if err != nil {
		return false
	}
	err = os.Chdir(directory)
	if err != nil {
		return false
	}
	defer func(dir string) {
		_ = os.Chdir(dir)
	}(currentDir)

	branch, errBranch := git.CurrentBranchName()
	if errBranch != nil {
		return false
	}

	if !slices.Contains(branchesOfInterest, branch) {
		return false
	}

	behindRemote, errBehindRemote := git.IsBehindRemote(remote, branch)
	if !behindRemote || errBehindRemote != nil {
		return false
	}

	return true
}
