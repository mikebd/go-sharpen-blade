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

func (b *board) canWin(player rune) bool {
	// Horizontal Win?
	for row := 0; row < rowSize; row++ {
		countPlayer := 0
		for col := 0; col < rowSize; col++ {
			if b.cells[row][col] == player {
				countPlayer++
			}
		}
		if countPlayer == rowSize {
			return true
		}
	}

	// Vertical Win?
	for col := 0; col < rowSize; col++ {
		countPlayer := 0
		for row := 0; row < rowSize; row++ {
			if b.cells[row][col] == player {
				countPlayer++
			}
		}
		if countPlayer == rowSize {
			return true
		}
	}

	// Diagonal Win from top-left?
	{
		countPlayer := 0
		for row := 0; row < rowSize; row++ {
			if b.cells[row][row] == player {
				countPlayer++
			}
		}
		if countPlayer == rowSize {
			return true
		}
	}

	// Diagonal Win from bottom-left?
	{
		countPlayer := 0
		for row := 0; row < rowSize; row++ {
			if b.cells[row][rowSize-1-row] == player {
				countPlayer++
			}
		}
		if countPlayer == rowSize {
			return true
		}
	}

	return false
}

func (b *board) evaluate() string {
	winO, winX := b.canWin(boardO), b.canWin(boardX)

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
