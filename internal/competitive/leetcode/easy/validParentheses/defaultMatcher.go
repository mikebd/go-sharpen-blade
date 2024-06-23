package validParentheses

import "strings"

const (
	// order is important, the index of each element in these strings must
	// match the index of its corresponding element in the other string
	opening = "({["
	closing = ")}]"
)

type defaultMatcher struct{}

func (defaultMatcher) IsOpening(r rune) bool {
	return strings.Contains(opening, string(r))
}

func (defaultMatcher) IsClosing(r rune) bool {
	return strings.Contains(closing, string(r))
}

func (defaultMatcher) IsMatch(openingRune, closingRune rune) bool {
	openingIndex := strings.Index(opening, string(openingRune))
	closingIndex := strings.Index(closing, string(closingRune))
	return openingIndex >= 0 && closingIndex >= 0 && openingIndex == closingIndex
}
