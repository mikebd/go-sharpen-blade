package board

// create creates a board with the given initial values.
func create(initialValues [size]string) Board {
	board := Board{}

	for row := 0; row < size; row++ {
		rowValues := initialValues[row]
		board.rows[row] = make(cellPointers, size)
		for col := 0; col < size; col++ {
			cellValue := emptyInitialValue
			if col < len(rowValues) {
				cellValue = rune(rowValues[col])
			}

			val := empty
			switch cellValue {
			case oneInitialValue:
				val = one
			case twoInitialValue:
				val = two
			case threeInitialValue:
				val = three
			case fourInitialValue:
				val = four
			case fiveInitialValue:
				val = five
			case sixInitialValue:
				val = six
			case sevenInitialValue:
				val = seven
			case eightInitialValue:
				val = eight
			case nineInitialValue:
				val = nine
			}
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

	return board
}
