package canBeMadeValidParentheses

type matcher struct{}

var theMatcher matcher

const (
	opening = '('
	closing = ')'
)

func (m matcher) IsOpening(r rune) bool {
	return r == opening
}

func (m matcher) IsClosing(r rune) bool {
	return r == closing
}

func (m matcher) IsMatch(openingRune, closingRune rune) bool {
	return m.IsOpening(openingRune) && m.IsClosing(closingRune)
}
