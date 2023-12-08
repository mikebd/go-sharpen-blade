package gitPull

import (
	"github.com/mikebd/go-util/pkg/directory"
	"github.com/mikebd/go-util/pkg/git"
	"slices"
)

var branchesOfInterest = []string{"main", "master"}

// remote is the repository short name that we want to track.
// This might become configurable in the future.
const remote = "origin"

// includeGitRepositoryDirectoryCheap returns true if the specified directory
// should be included in the list of directories to pull.  A directory
// is included if its branch is one of the branchesOfInterest and if it
// is behind the remote.
// This function includes only cheap tests that do not require goroutines.
func includeGitRepositoryDirectoryCheap(dir string) bool {
	_, restoreDir, err := directory.ChangeDirectory(dir)
	if err != nil {
		return false
	}
	defer restoreDir()

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
