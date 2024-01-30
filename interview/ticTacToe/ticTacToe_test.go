package ticTacToe

import "testing"

func Test_board_loadSuccess(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantBoard board
	}{
		{
			name:  "full board",
			input: "x,o,x,o,x,o,x,o,x",
			wantBoard: board{
				{boardX, boardO, boardX},
				{boardO, boardX, boardO},
				{boardX, boardO, boardX},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board{}
			if got := b.load(tt.input); got != "" {
				t.Errorf("load() = '%v', want no error", got)
			}
		})
	}
}

func Test_board_loadError(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "too many O 1",
			input: "o,o,,,,,,,",
		},
		{
			name:  "too many O 2",
			input: "o,x,o,o,,,,,",
		},
		{
			name:  "too many X",
			input: ",,x,x,o,,,x,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board{}
			if got := b.load(tt.input); got != resultError {
				t.Errorf("load() = '%v', want '%v'", got, resultError)
			}
		})
	}
}

func Test_board_loadInsufficientData(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "explicit empty board 1",
			input: ",,,,,,,,",
		},
		{
			name:  "explicit empty board 2",
			input: " , , , , , , , , ",
		},
		{
			name:  "implicit empty board",
			input: "",
		},
		{
			name:  "one cell",
			input: "x,,,,,,,,",
		},
		{
			name:  "two cells",
			input: "x,o,,,,,,,",
		},
		{
			name:  "three cells",
			input: "x,o,x,,,,,,",
		},
		{
			name:  "four cells",
			input: "x,o,x,o,,,,,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board{}
			if got := b.load(tt.input); got != resultInsufficientData {
				t.Errorf("load() = '%v', want '%v'", got, resultInsufficientData)
			}
		})
	}
}
