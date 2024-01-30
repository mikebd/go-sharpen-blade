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
				cells: [3][3]rune{
					{boardX, boardO, boardX},
					{boardO, boardX, boardO},
					{boardX, boardO, boardX},
				},
				countO: 4,
				countX: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board{}
			if got := b.load(tt.input); got != "" {
				t.Errorf("load() = '%v', want no error", got)
			}
			if b != tt.wantBoard {
				t.Errorf("load() = '%v', want '%v'", b, tt.wantBoard)
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
		name       string
		input      string
		wantCountO int
		wantCountX int
	}{
		{
			name:       "explicit empty board 1",
			input:      ",,,,,,,,",
			wantCountO: 0,
			wantCountX: 0,
		},
		{
			name:       "explicit empty board 2",
			input:      " , , , , , , , , ",
			wantCountO: 0,
			wantCountX: 0,
		},
		{
			name:       "implicit empty board",
			input:      "",
			wantCountO: 0,
			wantCountX: 0,
		},
		{
			name:       "one cell",
			input:      "x,,,,,,,,",
			wantCountO: 0,
			wantCountX: 1,
		},
		{
			name:       "two cells",
			input:      "x,o,,,,,,,",
			wantCountO: 1,
			wantCountX: 1,
		},
		{
			name:       "three cells",
			input:      "x,o,x,,,,,,",
			wantCountO: 1,
			wantCountX: 2,
		},
		{
			name:       "four cells",
			input:      "x,o,x,o,,,,,",
			wantCountO: 2,
			wantCountX: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board{}
			if got := b.load(tt.input); got != resultInsufficientData {
				t.Errorf("load() = '%v', want '%v'", got, resultInsufficientData)
			}
			if b.countO != tt.wantCountO {
				t.Errorf("load() wrong countO = '%v', want '%v'", b.countO, tt.wantCountO)
			}
			if b.countX != tt.wantCountX {
				t.Errorf("load() wrong countX = '%v', want '%v'", b.countX, tt.wantCountX)
			}
		})
	}
}
