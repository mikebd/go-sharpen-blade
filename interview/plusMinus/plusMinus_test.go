package plusMinus

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one+two", args{"one+two"}, 3},
		{"one+two-one-one+two+one", args{"one+two-one-one+two+one"}, 4},
		{"two-two-one-two", args{"two-two-one-two"}, -3},
		{"two", args{"two"}, 2},
		{"one-two", args{"one-two"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
