package board

import "fmt"

// createEmptyBoard creates a board with no initial values.
func createEmptyBoard() *Board {
	board, _ := create([size]string{})
	return board
}

// create creates a board with the given initial values.
// The initial values are given as an array strings that should
// not be longer than size.  emptyInitialValue is used to right
// pad if an initial value string is shorter than size.
func create(initialValues [size]string) (*Board, error) {
	board := Board{}

	for row := 0; row < size; row++ {
		rowValues := initialValues[row]

		if len(rowValues) > size {
			return nil, fmt.Errorf("invalid args to Board.create(), initialValues[%d] max length is %d, but it is of length %d",
				row, size, len(rowValues))
		}

		board.rows[row] = make(cellPointers, size)
		for col := 0; col < size; col++ {
			cellValue := emptyRuneValue
			if col < len(rowValues) {
				cellValue = rune(rowValues[col])
			}
			val := runeToValue(cellValue)
			board.rows[row][col] = &cell{row, col, val, val}
		}
	}

	for sct := 0; sct < size; sct++ {
		board.scts[sct] = make(cellPointers, size)
	}

	for col := 0; col < size; col++ {
		board.cols[col] = make(cellPointers, size)
		for row := 0; row < size; row++ {
			cellPtr := board.rows[row][col]
			board.cols[col][row] = cellPtr
			board.section(row, col)[offsetWithinSection(row, col)] = cellPtr
		}
	}

	return &board, nil
}
