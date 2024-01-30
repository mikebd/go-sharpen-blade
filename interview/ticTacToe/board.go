package ticTacToe

import "strings"

const (
	boardEmpty rune = 0
	boardO     rune = 'O'
	boardX     rune = 'X'

	fewestTurnsToWin int = 5
)

type board struct {
	cells  [3][3]rune
	countO int
	countX int
}

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

		b.cells[i/3][i%3] = boardRune
		if boardRune == boardO {
			b.countO++
		}
		if boardRune == boardX {
			b.countX++
		}
	}

	if b.countO > b.countX+1 || b.countX > b.countO+1 {
		return resultError
	}

	if b.countO+b.countX < fewestTurnsToWin {
		return resultInsufficientData
	}

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
