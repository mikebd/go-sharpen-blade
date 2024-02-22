package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
)

// createEmptyBoard creates a board with no initial values.
func createEmptyBoard() *board {
	board, _ := create([CellCount]string{})
	return board
}

// create creates a board with the given initial values.
// The initial values are given as an array strings that should
// not be longer than CellCount.  emptyInitialValue is used to right
// pad if an initial value string is shorter than CellCount.
func create(initialValues [CellCount]string) (*board, error) {
	board := board{}

	for row := 0; row < CellCount; row++ {
		rowValues := initialValues[row]

		if len(rowValues) > CellCount {
			return nil, fmt.Errorf("invalid args to board.create(), initialValues[%d] max length is %d, but it is of length %d",
				row, CellCount, len(rowValues))
		}

		board.rows[row] = make(CellPointers, CellCount)
		for col := 0; col < CellCount; col++ {
			cellValue := EmptyRuneValue
			if col < len(rowValues) {
				cellValue = rune(rowValues[col])
			}
			board.rows[row][col] = NewCell(row, col, cellValue)
		}
	}

	for sct := 0; sct < CellCount; sct++ {
		board.scts[sct] = make(CellPointers, CellCount)
	}

	for col := 0; col < CellCount; col++ {
		board.cols[col] = make(CellPointers, CellCount)
		for row := 0; row < CellCount; row++ {
			cellPtr := board.rows[row][col]
			board.cols[col][row] = cellPtr
			board.section(row, col)[offsetWithinSection(row, col)] = cellPtr
		}
	}

	return &board, nil
}
