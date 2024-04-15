package validParentheses

type Matcher interface {
	IsOpening(r rune) bool
	IsClosing(r rune) bool
	IsMatch(openingRune, closingRune rune) bool
}
