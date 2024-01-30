package ticTacToe

import "testing"

func Test_board_loadSuccess(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantBoard board
	}{
		{
			name:  "explicit empty board 1",
			input: ",,,,,,,,",
			wantBoard: board{
				{boardEmpty, boardEmpty, boardEmpty},
				{boardEmpty, boardEmpty, boardEmpty},
				{boardEmpty, boardEmpty, boardEmpty},
			},
		},
		{
			name:  "explicit empty board 2",
			input: " , , , , , , , , ",
			wantBoard: board{
				{boardEmpty, boardEmpty, boardEmpty},
				{boardEmpty, boardEmpty, boardEmpty},
				{boardEmpty, boardEmpty, boardEmpty},
			},
		},
		{
			name:  "implicit empty board",
			input: "",
			wantBoard: board{
				{boardEmpty, boardEmpty, boardEmpty},
				{boardEmpty, boardEmpty, boardEmpty},
				{boardEmpty, boardEmpty, boardEmpty},
			},
		},
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
