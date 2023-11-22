package gitPull

import (
	"fmt"
	"go-sharpen-blade/command"
)

func Register() {
	command.Register("git-pull", gitPull)
}

func gitPull() error {
	fmt.Println("Running git-pull")
	return nil
}
