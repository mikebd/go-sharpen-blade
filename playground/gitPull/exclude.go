package gitPull

import (
	"github.com/mikebd/go-util/pkg/git"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// remote is the repository short name that we want to track.
// This might become configurable in the future.
const remote = "origin"

type excludeDir struct {
	dir     string
	exclude bool
}

// excludeGitRepositoryDirectoryExpensive returns true if the specified directory
// should be excluded from the list of directories to pull.  i.e. if it is not
// behind the remote.
func excludeGitRepositoryDirectoryExpensive(
	parentDirectory string,
	excludeDirs []excludeDir,
	index int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	dir := excludeDirs[index].dir
	err := os.Chdir(filepath.Join(parentDirectory, dir))
	if err != nil {
		log.Println("[errChdir] Done with", dir, err)
		return
	}
	log.Println("chdir", dir)

	branch, errBranch := git.CurrentBranchName()
	if errBranch != nil {
		log.Println("[errBranch] Done with", dir, errBranch)
		return
	}
	log.Println("branch", dir)

	behindRemote, errBehindRemote := git.IsBehindRemote(remote, branch)
	if !behindRemote || errBehindRemote != nil {
		log.Println("[!behind || errBehind] Done with", dir)
		return
	}
	log.Println("behind", dir)

	log.Println("[end] Done with", dir)
	excludeDirs[index].exclude = false
}
