package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
)

// size is the number of cells in a row, column, or section.
// The square root of size must be an integer.
// There must be a constant of type value for each value from 1 to size.
const size int = 9

// cellPointers is a slice of pointers to cells.  It allows navigation of
// the Board by multiple dimensions (e.g. row, column, section).
// An update to a cell in one dimension will be reflected in all dimensions.
type cellPointers []*Cell

type cellPointersArray [size]cellPointers

// Board has cell accessors
type Board struct {
	// rows is an array of horizontal rows of cells, the 0th cell is the leftmost cell.
	rows cellPointersArray

	// cols is an array of vertical columns of cells, the 0th cell is the topmost cell.
	cols cellPointersArray

	// scts is an array of squares of cells, the 0th cell is the top left cell.
	// The 1th cell is the cell to the right of the 0th cell.
	// The (size-1)th cell is the bottom right cell.
	scts cellPointersArray
}

// cell returns a pointer to the cell.Cell at the given row and column.
func (board *Board) cell(row, col int) (*Cell, error) {
	if row < 0 || row >= size || col < 0 || col >= size {
		return nil, fmt.Errorf("board.cell() row and col must be between 0 and %d. row=%d, col=%d", size-1, row, col)
	}

	return board.rows[row][col], nil
}

// Size returns the number of cells in a row, column, or section.
func (board *Board) Size() int {
	return len(board.rows)
}
