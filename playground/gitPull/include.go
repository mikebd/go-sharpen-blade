package gitPull

import (
	"os"
)

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

	return true
}
