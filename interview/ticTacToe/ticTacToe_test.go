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
			name:  "O Win",
			input: "x,O,x,o,o,x,,o,",
			want:  "O Win",
		},
		{
			name:  "X Win",
			input: "x,X,x,o,o",
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
