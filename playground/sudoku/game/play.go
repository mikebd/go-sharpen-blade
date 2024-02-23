package game

import (
	. "go-sharpen-blade/playground/sudoku/cell"
)

func (game *game) play() error {
	for i := 0; i < 9; i++ {
		game.renderer.Render(game.turns, game.engine.Sections())
		turn, err := game.engine.MakeTurn(i+1, i, i, OneRuneValue)
		if err != nil {
			return err
		}
		game.turns = append(game.turns, &turn)
	}
	return nil
}
