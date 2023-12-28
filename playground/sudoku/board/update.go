package board

import (
	"fmt"
)

func (board *board) setValue(row, col int, value value) error {
	if row < 0 || row >= size || col < 0 || col >= size {
		return fmt.Errorf("board.setValue() row and col must be between 0 and %d. row=%d, col=%d", size-1, row, col)
	}

	board.rows[row][col].value = value
	return nil
}
