package plusMinus

import "testing"

func TestPlusMinus(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one+two", args{"one+two"}, 3},
		{
			"one+two+three+four+five+six+seven+eight+nine+zero",
			args{"one+two+three+four+five+six+seven+eight+nine+zero"},
			45},
		{"one+two-one-one+two+one", args{"one+two-one-one+two+one"}, 4},
		{"two-two-one-two", args{"two-two-one-two"}, -3},
		{"one-two", args{"one-two"}, -1},
		{"one", args{"one"}, 1},
		{"two", args{"two"}, 2},
		{"three", args{"three"}, 3},
		{"four", args{"four"}, 4},
		{"five", args{"five"}, 5},
		{"six", args{"six"}, 6},
		{"seven", args{"seven"}, 7},
		{"eight", args{"eight"}, 8},
		{"nine", args{"nine"}, 9},
		{"zero", args{"zero"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusMinus(tt.args.expr); got != tt.want {
				t.Errorf("plusMinus() = %v, want %v", got, tt.want)
			}
		})
	}
}
