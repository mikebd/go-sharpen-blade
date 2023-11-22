package gitPull

import "os"

type includeGitRepositoryDirectoryFunc func(directory string) bool

// findGitRepositoryDirectories returns a list of directories that contain a Git repository
func findGitRepositoryDirectories(parentDirectory string, include includeGitRepositoryDirectoryFunc) ([]string, error) {
	var result []string

	dirEntries, err := os.ReadDir(parentDirectory)
	if err != nil {
		return nil, err
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			dirName := parentDirectory + "/" + dirEntry.Name()
			_, err := os.Stat(dirName + "/.git")
			if err == nil {
				if include == nil || include(dirName) {
					result = append(result, dirName)
				}
			} else {
				subDirectories, err := findGitRepositoryDirectories(dirName, include)
				if err != nil {
					return nil, err
				}
				result = append(result, subDirectories...)
			}
		}
	}

	return result, nil
}
