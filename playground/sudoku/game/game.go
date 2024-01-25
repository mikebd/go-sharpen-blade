package game

import (
	"go-sharpen-blade/command"
	. "go-sharpen-blade/playground/sudoku/board"
	"go-sharpen-blade/playground/sudoku/game/port"
)

type game struct {
	board *Board
	turns []*port.Turn
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
