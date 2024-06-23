package canBeMadeValidParentheses

// flip returns a string with the rune at offset i flipped
// from opening to closing or closing to opening
func flip(s string, i int) string {
	if i < 0 || i >= len(s) {
		return s
	}

	prefix := s[0:i]
	r := rune(s[i])
	suffix := s[i+1:]

	switch {
	case theMatcher.IsOpening(r):
		return prefix + string(closing) + suffix
	case theMatcher.IsClosing(r):
		return prefix + string(opening) + suffix
	default:
		return s
	}
}
