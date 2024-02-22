package game

import (
	"go-sharpen-blade/command"
	. "go-sharpen-blade/playground/sudoku/board"
	"go-sharpen-blade/playground/sudoku/game/port"
	"go-sharpen-blade/playground/sudoku/ui/cli"
)

type game struct {
	engine   port.Engine
	renderer port.Renderer
	turns    []*port.Turn
}

func Register() {
	command.Register("sudoku", newGame)
}

func newGame() error {
	game := game{
		engine:   GenerateBoard(),
		renderer: &cli.Console{},
	}

	return game.play()
}
