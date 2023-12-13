package gitPull

import (
	"github.com/mikebd/go-util/pkg/git"
	"log"
	"os"
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
// behind the remote.  Assumes excludeDirs[index].exclude is true, sets it to false
// if the repository is behind the remote.
func excludeGitRepositoryDirectoryExpensive(
	parentDirectory string,
	excludeDirs []excludeDir,
	index int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	pwd, _ := os.Getwd()
	log.Println("pwd", pwd)

	dir := excludeDirs[index].dir
	globalOptions := git.GlobalOptions{AsIfIn: dir}

	branch, errBranch := git.CurrentBranchName(globalOptions)
	if errBranch != nil {
		log.Println("[errBranch] Done with", dir, errBranch)
		return
	}

	errFetch := git.Fetch(remote, branch, globalOptions)
	if errFetch != nil {
		log.Println("[errFetch] Done with", dir, errFetch)
		return
	}

	behindRemote, errBehindRemote := git.IsBehindRemote(remote, branch, globalOptions)
	if !behindRemote || errBehindRemote != nil {
		log.Println("[!behind || errBehind] Done with", dir, "behindRemote", behindRemote, "errBehindRemote", errBehindRemote)
		return
	}
	log.Println("behind", dir)

	excludeDirs[index].exclude = false
}
