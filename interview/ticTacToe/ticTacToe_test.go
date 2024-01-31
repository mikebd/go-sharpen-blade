package ticTacToe

import "testing"

func Test_ticTacToe(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Tie",
			input: "x,o,x,x,o,x,o,x,o",
			want:  "Tie",
		},
		{
			name:  "O Win - vertical",
			input: "x, O,x,o,o,x,,o,",
			want:  "O Win",
		},
		{
			name:  "X Win - horizontal",
			input: "x,X,x,o,o",
			want:  "X Win",
		},
		{
			name:  "X Win - diagonal top-left to bottom-right",
			input: "x,o,x,o,x,o,,,x",
			want:  "X Win",
		},
		{
			name:  "X Win - diagonal bottom-left to top-right",
			input: "x,o,x,o,x,o,x,,",
			want:  "X Win",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ticTacToe(tt.input); got != tt.want {
				t.Errorf("ticTacToe() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
