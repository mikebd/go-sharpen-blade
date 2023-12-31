package gitPull

import (
	"os"
	"path/filepath"
)

// findGitRepositoryDirectories returns a list of directories that contain a Git repository
func findGitRepositoryDirectories(parentDirectory string) ([]string, error) {
	result, err := findGitRepositoryDirectoriesCheap(parentDirectory)

	if err != nil {
		return nil, err
	}

	return excludeGitRepositoryDirectoriesExpensive(result), nil
}

func findGitRepositoryDirectoriesCheap(parentDirectory string) ([]string, error) {
	var result []string

	dirEntries, err := os.ReadDir(parentDirectory)
	if err != nil {
		return nil, err
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			dirName := filepath.Join(parentDirectory, dirEntry.Name())
			_, err := os.Stat(filepath.Join(dirName, ".git"))
			if err == nil {
				if includeGitRepositoryDirectoryCheap(dirName) {
					result = append(result, dirName)
				}
			} else {
				subDirectories, err := findGitRepositoryDirectoriesCheap(dirName)
				if err != nil {
					return nil, err
				}
				result = append(result, subDirectories...)
			}
		}
	}

	return result, nil
}

func excludeGitRepositoryDirectoriesExpensive(directories []string) []string {
	return directories
}
