package board

import (
	"fmt"
)

// setValue sets the value of the cell at the given row and column.
// It returns the previous value of the cell.
// It returns an error if the row or column is out of range.
func (board *Board) setValue(row, col int, value value) (value, error) {
	if row < 0 || row >= size || col < 0 || col >= size {
		return nilValue, fmt.Errorf("Board.setValue() row and col must be between 0 and %d. row=%d, col=%d", size-1, row, col)
	}

	previousValue := board.rows[row][col].value
	board.rows[row][col].value = value
	return previousValue, nil
}
