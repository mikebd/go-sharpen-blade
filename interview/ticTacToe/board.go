package ticTacToe

import "strings"

const (
	boardEmpty rune = 0
	boardO     rune = 'O'
	boardX     rune = 'X'

	rowSize          int = 3
	boardSize        int = rowSize * rowSize
	fewestTurnsToWin int = rowSize*2 - 1
)

type board struct {
	cells  [rowSize][rowSize]rune
	countO int
	countX int
}

func (b *board) hasWon(player rune) bool {
	// Player piece count per diagonal
	countPlayerTopLeft, countPlayerBottomLeft := 0, 0

	for i := 0; i < rowSize; i++ {
		// Diagonal win?
		if b.cells[i][i] == player {
			countPlayerTopLeft++
		}
		if b.cells[i][rowSize-1-i] == player {
			countPlayerBottomLeft++
		}
		if i == rowSize-1 && (countPlayerTopLeft == rowSize || countPlayerBottomLeft == rowSize) {
			return true
		}

		// Horizontal or vertical win?
		countPlayerHorizontal, countPlayerVertical := 0, 0
		for j := 0; j < rowSize; j++ {
			if b.cells[i][j] == player {
				countPlayerHorizontal++
			}
			if b.cells[j][i] == player {
				countPlayerVertical++
			}
		}
		if countPlayerHorizontal == rowSize || countPlayerVertical == rowSize {
			return true
		}
	}

	return false
}

func (b *board) evaluate() string {
	// For extreme optimization, hasWon() could be modified to test both players at once
	winO, winX := b.hasWon(boardO), b.hasWon(boardX)

	switch {
	case winO && winX:
		return resultError
	case winO:
		return resultOWin
	case winX:
		return resultXWin
	case b.countO+b.countX == boardSize:
		return resultTie
	default:
		return resultInsufficientData
	}
}

// loadBoard takes a string representing a tic-tac-toe board and loads it into the board.
// The return value is an empty string if the board was loaded successfully.  Otherwise, it is a result... constant.
func (b *board) load(input string) string {
	cells := strings.Split(input, ",")
	if len(cells) > boardSize {
		return resultError
	}

	for i, cellString := range cells {
		boardRune, errorString := b.cellStringToBoardRune(cellString)
		if len(errorString) > 0 {
			return errorString
		}

		b.cells[i/rowSize][i%rowSize] = boardRune
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
	cell = strings.TrimSpace(cell)
	cellLen := len(cell)
	if cellLen == 0 {
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
