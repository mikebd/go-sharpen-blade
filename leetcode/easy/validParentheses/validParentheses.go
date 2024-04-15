// https://leetcode.com/problems/valid-parentheses

package validParentheses

import (
	"github.com/mikebd/go-util/pkg/collection"
	"strings"
)

const (
	// order is important, the index of each element in these strings must
	// match the index of its corresponding element in the other string
	opening = "({["
	closing = ")}]"
)

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	stack := collection.NewStackWithCapacity[rune](len(s))

	for _, r := range s {
		if stack.IsEmpty() {
			if isOpening(r) {
				stack.Push(r)
			} else {
				return false
			}
		} else {
			switch {
			case isOpening(r):
				stack.Push(r)
			case isClosing(r):
				prevOpening, ok := stack.Pop()
				if !ok || !isMatch(prevOpening, r) {
					return false
				}
			}
		}
	}

	return stack.IsEmpty()
}

func isOpening(r rune) bool {
	return strings.Contains(opening, string(r))
}

func isClosing(r rune) bool {
	return strings.Contains(closing, string(r))
}

func isMatch(openingRune, closingRune rune) bool {
	openingIndex := strings.Index(opening, string(openingRune))
	closingIndex := strings.Index(closing, string(closingRune))
	return openingIndex >= 0 && closingIndex >= 0 && openingIndex == closingIndex
}
