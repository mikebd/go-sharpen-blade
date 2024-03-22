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
		{"one+two-one-one+two+one", args{"one+two-one-one+two+one"}, 4},
		{"two-two-one-two", args{"two-two-one-two"}, -3},
		{"two", args{"two"}, 2},
		{"one-two", args{"one-two"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusMinus(tt.args.expr); got != tt.want {
				t.Errorf("plusMinus() = %v, want %v", got, tt.want)
			}
		})
	}
}
