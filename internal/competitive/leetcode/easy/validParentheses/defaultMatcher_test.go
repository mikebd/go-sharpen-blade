package validParentheses

import (
	"fmt"
	"strings"
)

func Example_isOpening() {
	m := defaultMatcher{}

	fmt.Println(m.IsOpening('('))
	fmt.Println(m.IsOpening('{'))
	fmt.Println(m.IsOpening('['))

	fmt.Println(m.IsOpening(' '))
	fmt.Println(m.IsOpening(')'))
	fmt.Println(m.IsOpening('}'))
	fmt.Println(m.IsOpening(']'))

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
	m := defaultMatcher{}

	fmt.Println(m.IsClosing(')'))
	fmt.Println(m.IsClosing('}'))
	fmt.Println(m.IsClosing(']'))

	fmt.Println(m.IsClosing(' '))
	fmt.Println(m.IsClosing('('))
	fmt.Println(m.IsClosing('{'))
	fmt.Println(m.IsClosing('['))

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
	m := defaultMatcher{}

	fmt.Println(m.IsMatch('(', ')'))
	fmt.Println(m.IsMatch('{', '}'))
	fmt.Println(m.IsMatch('[', ']'))

	fmt.Println(m.IsMatch('(', '}'))
	fmt.Println(m.IsMatch('(', ']'))
	fmt.Println(m.IsMatch('{', ')'))
	fmt.Println(m.IsMatch('{', ']'))
	fmt.Println(m.IsMatch('[', ')'))
	fmt.Println(m.IsMatch('[', '}'))

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
