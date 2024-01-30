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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board{}
			if got := b.load(tt.input); got != "" {
				t.Errorf("load() = %v, want no error", got)
			}
		})
	}
}
