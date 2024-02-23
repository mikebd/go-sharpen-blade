package port

import . "go-sharpen-blade/playground/sudoku/cell"

// Engine is an interface for making turns and accessing the sections of a board
type Engine interface {
	MakeTurn(turnNumber, row, col int, value rune) (Turn, error)

	Sections() CellPointersArray
}
