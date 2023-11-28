package gitPull

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type includeGitRepositoryDirectoryFunc func(directory string) bool

// findGitRepositoryDirectories returns a list of directories that contain a Git repository
func findGitRepositoryDirectories(
	parentDirectory string,
	include includeGitRepositoryDirectoryFunc,
) ([]string, error) {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go findGitRepositoryDirectoriesParallel(parentDirectory, include, &wg, ch)

	log.Println("Waiting for goroutines to finish")
	wg.Wait()
	log.Println("All goroutines finished")

	var result []string
	for i := 0; i < len(ch); i++ {
		result = append(result, <-ch)
	}

	log.Println("Found", len(result), "Git repositories", result)
	return result, nil
}

var logDiagnostics = true
var count uint32

func findGitRepositoryDirectoriesParallel(
	parentDirectory string,
	include includeGitRepositoryDirectoryFunc,
	wg *sync.WaitGroup,
	ch chan string,
) {
	var counter uint32
	if logDiagnostics {
		counter = atomic.AddUint32(&count, 1)
		log.Printf("Entering #%04d %s\n", counter, parentDirectory)
	}
	defer func() {
		if logDiagnostics {
			log.Printf("Leaving  #%04d %s\n", counter, parentDirectory)
		}
		wg.Done()
	}()

	dirEntries, err := os.ReadDir(parentDirectory)
	if err != nil {
		return
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			dirName := filepath.Join(parentDirectory, dirEntry.Name())
			_, err := os.Stat(filepath.Join(dirName, ".git"))
			if err == nil {
				if include == nil || include(dirName) {
					ch <- dirName
				}
			} else {
				wg.Add(1)
				go findGitRepositoryDirectoriesParallel(dirName, include, wg, ch)
			}
		}
	}
}
