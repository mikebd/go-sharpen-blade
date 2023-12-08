package gitPull

import (
	"go-sharpen-blade/command"
	"log"
)

func Register() {
	command.Register("git-pull", gitPull)
}

func gitPull() error {
	directories, err := findGitRepositoryDirectories(".")
	if err != nil {
		return err
	}

	for _, directory := range directories {
		log.Println("Pulling", directory)
	}

	return nil
}
