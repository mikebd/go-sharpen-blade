package gitPull

import (
	"os"
	"os/exec"
	"strconv"
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

	if !isBranchOfInterest(branch) {
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

func isBranchOfInterest(branch string) bool {
	for _, branchOfInterest := range branchesOfInterest {
		if branch == branchOfInterest {
			return true
		}
	}
	return false
}

func isBehindRemote(branch string) bool {
	fetchErr := fetch(branch)
	if fetchErr != nil {
		return false
	}

	cmd := exec.Command("git", "rev-list", "--count", branch+".."+remote+"/"+branch)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return isCommandOutputGreaterThanZero(output)
}

func fetch(branch string) error {
	cmd := exec.Command("git", "fetch", remote, branch)
	return cmd.Run()
}

func isCommandOutputGreaterThanZero(output []byte) bool {
	if len(output) <= 1 {
		return false
	}
	value, err := strconv.Atoi(string(output)[:len(output)-1])
	if err != nil {

	}
	return value > 0
}
