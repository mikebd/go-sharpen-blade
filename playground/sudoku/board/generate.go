package board

import game "go-sharpen-blade/playground/sudoku/game/port"

func GenerateBoard() game.Engine {
	return createEmptyBoard()
}
