package board

import (
	game "go-sharpen-blade/playground/sudoku/game/turn"
)

type turn struct {
	number, row, col     int
	valid                bool
	previousValue, value rune
}

// MakeTurn is an adaptor that exposes the internal setValue method as a Turn
func (board *Board) MakeTurn(turnNumber, row, col int, turnValueRune rune) (game.Turn, error) {
	previousValue, err := board.setValue(row, col, runeToValue(turnValueRune))
	if err != nil {
		return nil, err
	}
	invalid := board.Validate()

	return &turn{
		number:        turnNumber,
		row:           row,
		col:           col,
		valid:         !invalid.Any(),
		previousValue: previousValue.rune(),
		value:         turnValueRune,
	}, nil
}

func (turn *turn) Number() int {
	return turn.number
}

func (turn *turn) Valid() bool {
	return turn.valid
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
