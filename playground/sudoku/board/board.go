package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
)

// board has cell accessors and implements to game.Engine interface.
type board struct {
	// rows is an array of horizontal rows of cells, the 0th cell is the leftmost cell.
	rows CellPointersArray

	// cols is an array of vertical columns of cells, the 0th cell is the topmost cell.
	cols CellPointersArray

	// scts is an array of squares of cells, the 0th cell is the top left cell.
	// The 1th cell is the cell to the right of the 0th cell.
	// The (CellCount-1)th cell is the bottom right cell.
	scts CellPointersArray
}

func (b *board) Sections() CellPointersArray {
	return b.scts
}

// cell returns a pointer to the cell.Cell at the given row and column.
func (b *board) cell(row, col int) (*Cell, error) {
	if row < 0 || row >= CellCount || col < 0 || col >= CellCount {
		return nil, fmt.Errorf("board.cell() row and col must be between 0 and %d. row=%d, col=%d", CellCount-1, row, col)
	}

	return b.rows[row][col], nil
}

// Size returns the number of cells in a row, column, or section.
func (b *board) Size() int {
	return len(b.rows)
}
