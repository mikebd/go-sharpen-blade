package ticTacToe

const (
	resultError            = "Error"
	resultInsufficientData = "Insufficient Data"
	resultTie              = "Tie"
	resultWin              = "O Win"
	resultXWin             = "X Win"
)

// ticTacToe takes a string representing a tic-tac-toe board and returns a string representing the result of the game.
// The input string should be a comma separated list of characters representing the board.
// The characters are in row-major order. The characters are either ['x', 'X'], ['o', 'O'], or ' '.
func ticTacToe(input string) string {
	board := board{}

	result := board.load(input)
	if len(result) > 0 {
		return result
	}

	return result
}
