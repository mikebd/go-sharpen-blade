package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
)

// Board has cell accessors
type Board struct {
	// rows is an array of horizontal rows of cells, the 0th cell is the leftmost cell.
	rows CellPointersArray

	// cols is an array of vertical columns of cells, the 0th cell is the topmost cell.
	cols CellPointersArray

	// scts is an array of squares of cells, the 0th cell is the top left cell.
	// The 1th cell is the cell to the right of the 0th cell.
	// The (CellCount-1)th cell is the bottom right cell.
	scts CellPointersArray
}

// cell returns a pointer to the cell.Cell at the given row and column.
func (board *Board) cell(row, col int) (*Cell, error) {
	if row < 0 || row >= CellCount || col < 0 || col >= CellCount {
		return nil, fmt.Errorf("board.cell() row and col must be between 0 and %d. row=%d, col=%d", CellCount-1, row, col)
	}

	return board.rows[row][col], nil
}

// Size returns the number of cells in a row, column, or section.
func (board *Board) Size() int {
	return len(board.rows)
}
