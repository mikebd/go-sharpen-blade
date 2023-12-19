package elimination

func eliminate(board [][]uint8) [][]uint8 {
	result := make([][]uint8, len(board))
	for i := range board {
		boardRow := board[i]
		result[i] = make([]uint8, len(boardRow))
		copy(result[i], boardRow)

		twoBefore, oneBefore := uint8(0), uint8(0)
		for index, value := range boardRow {
			switch {
			case index == 0:
				twoBefore = value
			case index == 1:
				oneBefore = value
			default:
				if twoBefore == oneBefore && oneBefore == value {
					result[i][index-2] = 0
					result[i][index-1] = 0
					result[i][index] = 0
				}
				twoBefore = oneBefore
				oneBefore = value
			}
		}
	}

	return result
}
