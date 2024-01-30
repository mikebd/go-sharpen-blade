package ticTacToe

import "strings"

const (
	boardEmpty rune = 0
	boardO     rune = 'O'
	boardX     rune = 'X'
)

type board [3][3]rune

// loadBoard takes a string representing a tic-tac-toe board and loads it into the board.
// The return value is an empty string if the board was loaded successfully.  Otherwise, it is a result... constant.
func (b *board) load(input string) string {
	cells := strings.Split(input, ",")
	if len(cells) > 9 {
		return resultError
	}
	for i, cellString := range cells {
		boardRune, errorString := b.cellStringToBoardRune(cellString)
		if len(errorString) > 0 {
			return errorString
		}
		b[i/3][i%3] = boardRune
	}

	// TODO: Error if there are too many Xs or Os

	return ""
}

func (b *board) cellStringToBoardRune(cell string) (rune, string) {
	cellLen := len(cell)
	if cellLen == 0 || cellLen > 1 {
		return boardEmpty, ""
	}
	if cellLen > 1 {
		return boardEmpty, resultError
	}

	switch cell[0] {
	case 'x', 'X':
		return boardX, ""
	case 'o', 'O':
		return boardO, ""
	case ' ':
		return boardEmpty, ""
	default:
		return boardEmpty, resultError
	}
}
