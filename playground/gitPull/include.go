package gitPull

import (
	"os"
	"os/exec"
)

var branchesOfInterest = []string{"main", "master"}

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
