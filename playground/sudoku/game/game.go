package game

import (
	"go-sharpen-blade/command"
	. "go-sharpen-blade/playground/sudoku/board"
)

type game struct {
	board *Board
}

func Register() {
	command.Register("sudoku", newGame)
}

func newGame() error {
	game := &game{
		board: GenerateBoard(),
	}

	return game.play()
}
