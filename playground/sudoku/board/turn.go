package board

import (
	game "go-sharpen-blade/playground/sudoku/game/port"
)

type turn struct {
	number, row, col     int
	valid                bool
	previousValue, value rune
}

// MakeTurn is an adaptor that exposes the internal setValue method as a Turn
func (b *board) MakeTurn(turnNumber, row, col int, turnRuneValue rune) (game.Turn, error) {
	previousRuneValue, err := b.setValue(row, col, turnRuneValue)
	if err != nil {
		return nil, err
	}
	invalid := b.Validate()

	return &turn{
		number:        turnNumber,
		row:           row,
		col:           col,
		valid:         !invalid.Any(),
		previousValue: previousRuneValue,
		value:         turnRuneValue,
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
