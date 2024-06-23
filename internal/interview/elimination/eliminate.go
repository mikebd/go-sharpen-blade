package elimination

func eliminate(board [][]uint8) [][]uint8 {
	result := make([][]uint8, len(board))

	var twoAbove, oneAbove []uint8
	for rowIndex := range board {
		boardRow := board[rowIndex]
		result[rowIndex] = make([]uint8, len(boardRow))
		copy(result[rowIndex], boardRow)

		if rowIndex == 0 {
			twoAbove = make([]uint8, len(boardRow))
			oneAbove = make([]uint8, len(boardRow))
		}
		twoBefore, oneBefore := uint8(0), uint8(0)
		for colIndex, value := range boardRow {
			switch {
			case rowIndex == 0:
				twoAbove[colIndex] = value
			case rowIndex == 1:
				oneAbove[colIndex] = value
			default:
				if twoAbove[colIndex] == value && oneAbove[colIndex] == value {
					result[rowIndex-2][colIndex] = 0
					result[rowIndex-1][colIndex] = 0
					result[rowIndex][colIndex] = 0
				}
				twoAbove[colIndex] = oneAbove[colIndex]
				oneAbove[colIndex] = value
			}
			switch {
			case colIndex == 0:
				twoBefore = value
			case colIndex == 1:
				oneBefore = value
			default:
				if twoBefore == value && oneBefore == value {
					result[rowIndex][colIndex-2] = 0
					result[rowIndex][colIndex-1] = 0
					result[rowIndex][colIndex] = 0
				}
				twoBefore = oneBefore
				oneBefore = value
			}
		}
	}

	return result
}
