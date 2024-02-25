package game

import (
	. "go-sharpen-blade/playground/sudoku/cell"
)

func (game *game) play() error {
	defer func() {
		game.renderer.Render(game.turns, game.engine.Sections())
	}()

	for i := 0; i < 9; i++ {
		if !game.renderer.Overwrites() {
			game.renderer.Render(game.turns, game.engine.Sections())
		}
		turn, err := game.engine.MakeTurn(i+1, i, i, OneRuneValue+rune(i))
		if err != nil {
			return err
		}
		game.turns = append(game.turns, &turn)
	}

	return nil
}
