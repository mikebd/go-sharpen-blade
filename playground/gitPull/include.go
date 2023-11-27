package gitPull

import (
	"github.com/mikebd/go-util/pkg/git"
	"github.com/mikebd/go-util/pkg/shell"
	"os"
	"os/exec"
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

	branch := currentBranchName()

	if !slices.Contains(branchesOfInterest, branch) {
		return false
	}

	if !isBehindRemote(branch) {
		return false
	}

	return true
}

func currentBranchName() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil || len(output) <= 1 {
		return ""
	}
	return string(output)[:len(output)-1]
}

func isBehindRemote(branch string) bool {
	fetchErr := git.Fetch(remote, branch)
	if fetchErr != nil {
		return false
	}

	cmd := exec.Command("git", "rev-list", "--count", branch+".."+remote+"/"+branch)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return shell.IsOutputGreaterThanZero(output)
}
