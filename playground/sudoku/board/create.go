package board

const (
	emptyInitialValue rune = ' '
	oneInitialValue   rune = '1'
	twoInitialValue   rune = '2'
	threeInitialValue rune = '3'
	fourInitialValue  rune = '4'
	fiveInitialValue  rune = '5'
	sixInitialValue   rune = '6'
	sevenInitialValue rune = '7'
	eightInitialValue rune = '8'
	nineInitialValue  rune = '9'
)

func create(initialValues [size]string) board {
	board := board{}

	for rowOffset := 0; rowOffset < size; rowOffset++ {
		rowValues := initialValues[rowOffset]
		board.rows[rowOffset] = make(row, size)
		for colOffset := 0; colOffset < size; colOffset++ {
			cellValue := emptyInitialValue
			if colOffset < len(rowValues) {
				cellValue = rune(rowValues[colOffset])
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
			board.rows[rowOffset][colOffset] = cell{rowOffset, colOffset, val, val}
		}
	}

	for sctOffset := 0; sctOffset < size; sctOffset++ {
		board.scts[sctOffset] = make(sct, size)
	}

	for colOffset := 0; colOffset < size; colOffset++ {
		board.cols[colOffset] = make(col, size)
		for rowOffset := 0; rowOffset < size; rowOffset++ {
			cell := &board.rows[rowOffset][colOffset]
			board.cols[colOffset][rowOffset] = cell
			board.section(rowOffset, colOffset)[offsetWithinSection(rowOffset, colOffset)] = cell
		}
	}

	return board
}
