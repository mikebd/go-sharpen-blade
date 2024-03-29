package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
)

// setValue sets the rune value of the cell.Cell at the given row and column.
// It returns the previous rune value of the cell.
// It returns an error if the row or column is out of range.
func (b *board) setValue(row, col int, runeValue rune) (rune, error) {
	c, err := b.cell(row, col)
	if err != nil {
		return NilRuneValue, fmt.Errorf("board.setValue() row and col must be between 0 and %d. row=%d, col=%d", CellCount-1, row, col)
	}

	return c.SetValue(runeValue), nil
}
