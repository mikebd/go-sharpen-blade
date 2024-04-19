package canBeMadeValidParentheses

import (
	"strconv"
	"testing"
)

func Test_flip(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		args args
		want string
	}{
		{args{"(", 0}, ")"},
		{args{"((", 0}, ")("},
		{args{"((", 1}, "()"},
		{args{"(((", 0}, ")(("},
		{args{"(((", 1}, "()("},
		{args{"(((", 2}, "(()"},
		{args{")", 0}, "("},
		{args{"))", 0}, "()"},
		{args{"))", 1}, ")("},
		{args{")))", 0}, "())"},
		{args{")))", 1}, ")()"},
		{args{")))", 2}, "))("},
	}
	for _, tt := range tests {
		t.Run(tt.args.s+" "+strconv.Itoa(tt.args.i), func(t *testing.T) {
			if got := flip(tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("flip() = %v, want %v", got, tt.want)
			}
		})
	}
}
