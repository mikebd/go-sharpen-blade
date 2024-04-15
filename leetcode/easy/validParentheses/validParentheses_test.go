package validParentheses

import (
	"fmt"
	"strings"
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
			if got := isValid(tt.arg); !got {
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
			if got := isValid(tt.arg); got {
				t.Errorf("isValid() = %v, want: false", got)
			}
		})
	}
}

func Example_isOpening() {
	fmt.Println(isOpening('('))
	fmt.Println(isOpening('{'))
	fmt.Println(isOpening('['))

	fmt.Println(isOpening(' '))
	fmt.Println(isOpening(')'))
	fmt.Println(isOpening('}'))
	fmt.Println(isOpening(']'))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func Example_isClosing() {
	fmt.Println(isClosing(')'))
	fmt.Println(isClosing('}'))
	fmt.Println(isClosing(']'))

	fmt.Println(isClosing(' '))
	fmt.Println(isClosing('('))
	fmt.Println(isClosing('{'))
	fmt.Println(isClosing('['))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func Example_isMatchOpeningImpl() {
	fmt.Println(strings.Index(opening, "("))
	fmt.Println(strings.Index(opening, string('(')))
	fmt.Println(strings.Index(opening, "{"))
	fmt.Println(strings.Index(opening, string('{')))
	fmt.Println(strings.Index(opening, "["))
	fmt.Println(strings.Index(opening, string('[')))
	// Output:
	// 0
	// 0
	// 1
	// 1
	// 2
	// 2
}

func Example_isMatchClosingImpl() {
	fmt.Println(strings.Index(closing, ")"))
	fmt.Println(strings.Index(closing, string(')')))
	fmt.Println(strings.Index(closing, "}"))
	fmt.Println(strings.Index(closing, string('}')))
	fmt.Println(strings.Index(closing, "]"))
	fmt.Println(strings.Index(closing, string(']')))
	// Output:
	// 0
	// 0
	// 1
	// 1
	// 2
	// 2
}

func Example_isMatch() {
	fmt.Println(isMatch('(', ')'))
	fmt.Println(isMatch('{', '}'))
	fmt.Println(isMatch('[', ']'))

	fmt.Println(isMatch('(', '}'))
	fmt.Println(isMatch('(', ']'))
	fmt.Println(isMatch('{', ')'))
	fmt.Println(isMatch('{', ']'))
	fmt.Println(isMatch('[', ')'))
	fmt.Println(isMatch('[', '}'))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
}
