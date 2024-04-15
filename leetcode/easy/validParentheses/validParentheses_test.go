package validParentheses

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		arg string
	}{
		{"()"},
		{"()[]{}"},
		{"([]{})"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := isValidDefault(tt.arg); !got {
				t.Errorf("isValid() = %v, want: true", got)
			}
		})
	}
}

func Test_isNotValid(t *testing.T) {
	tests := []struct {
		arg string
	}{
		{""},
		{"(]"},
		{"()]"},
		{"[()"},
		{"]()"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := isValidDefault(tt.arg); got {
				t.Errorf("isValid() = %v, want: false", got)
			}
		})
	}
}
