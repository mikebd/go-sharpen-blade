package gitPull

import (
	"os"
	"path/filepath"
	"sync"
)

// findGitRepositoryDirectories returns a list of directories that contain a Git repository
func findGitRepositoryDirectories(parentDirectory string) ([]string, error) {
	result, err := findGitRepositoryDirectoriesCheap(parentDirectory)

	if err != nil {
		return nil, err
	}

	return excludeGitRepositoryDirectoriesExpensive(parentDirectory, result), nil
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
			fileInfo, err := os.Stat(filepath.Join(dirName, ".git"))
			// TODO: Handle git symbolic links to external git-dir
			if err == nil && fileInfo.IsDir() {
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

func excludeGitRepositoryDirectoriesExpensive(parentDirectory string, directories []string) []string {
	excludeDirs := make([]excludeDir, len(directories))
	var wg sync.WaitGroup

	wg.Add(len(directories))
	for i, directory := range directories {
		excludeDirs[i] = excludeDir{directory, true}
		go excludeGitRepositoryDirectoryExpensive(parentDirectory, excludeDirs, i, &wg)
	}

	wg.Wait()

	result := make([]string, 0, len(directories))

	for _, excludeDir := range excludeDirs {
		if !excludeDir.exclude {
			result = append(result, excludeDir.dir)
		}
	}

	return result
}
