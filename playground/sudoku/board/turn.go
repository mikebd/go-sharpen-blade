package board

import (
	game "go-sharpen-blade/playground/sudoku/game/turn"
)

type turn struct {
	number, row, col     int
	error                error
	invalid              *Invalid
	previousValue, value rune
}

// MakeTurn is an adaptor that exposes the internal setValue method as a Turn
func (board *Board) MakeTurn(turnNumber, row, col int, turnValueRune rune) game.Turn {
	previousValue, err := board.setValue(row, col, runeToValue(turnValueRune))
	invalid := board.Validate()

	return &turn{
		number:        turnNumber,
		row:           row,
		col:           col,
		error:         err,
		invalid:       invalid,
		previousValue: previousValue.rune(),
		value:         turnValueRune,
	}
}

func (turn *turn) Number() int {
	return turn.number
}

func (turn *turn) Error() error {
	return turn.error
}

func (turn *turn) Valid() bool {
	return turn.invalid.Any()
}

func (turn *turn) Row() int {
	return turn.row
}

func (turn *turn) Col() int {
	return turn.col
}

func (turn *turn) PreviousValue() rune {
	return turn.previousValue
}

func (turn *turn) Value() rune {
	return turn.value
}
