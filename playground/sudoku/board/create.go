package board

func create(initialValues [size]string) board {
	board := board{}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			board[row][col] = cell{row, col, empty, empty}
		}
	}

	return board
}
