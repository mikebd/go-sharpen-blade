// https://leetcode.com/problems/valid-parentheses

package validParentheses

import (
	"github.com/mikebd/go-util/pkg/collection"
)

func isValidDefault(s string) bool {
	return IsValid(s, defaultMatcher{})
}

func IsValid(s string, m Matcher) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	stack := collection.NewStackWithCapacity[rune](len(s))

	for _, r := range s {
		if stack.IsEmpty() {
			if m.IsOpening(r) {
				stack.Push(r)
			} else {
				return false
			}
		} else {
			switch {
			case m.IsOpening(r):
				stack.Push(r)
			case m.IsClosing(r):
				prevOpening, ok := stack.Pop()
				if !ok || !m.IsMatch(prevOpening, r) {
					return false
				}
			}
		}
	}

	return stack.IsEmpty()
}
