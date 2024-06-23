package elimination

// createBoard transforms a slice of strings into a 2D slice of uint8
// where each uint8 represents a character in the string
// e.g. ["135", "246", "789"] -> [[1 3 5] [2 4 6] [7 8 9]]
func createBoard(input []string) [][]uint8 {
	board := make([][]uint8, len(input))
	for i, row := range input {
		board[i] = make([]uint8, len(row))
		for j := range board[i] {
			board[i][j] = row[j] - '0'
		}
	}
	return board
}
