// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid

package canBeMadeValidParentheses

import (
	"go-sharpen-blade/internal/competitive/leetcode/easy/validParentheses"
)

func canBeValid(s string, locked string) bool {
	if len(s) == 0 || len(s)%2 != 0 || len(locked) != len(s) {
		return false
	}

	if validParentheses.IsValid(s, theMatcher) {
		return true
	}

	uo := unlockedOffsets(locked)

	for _, i := range uo {
		flippedOnce := flip(s, i)
		if validParentheses.IsValid(flippedOnce, theMatcher) {
			return true
		}

		for _, j := range uo {
			if i != j {
				flippedTwice := flip(flippedOnce, j)
				if validParentheses.IsValid(flippedTwice, theMatcher) {
					return true
				}
			}
		}
	}

	return false
}

func unlockedOffsets(locked string) []int {
	result := make([]int, 0, len(locked))

	for i, l := range locked {
		if l == '0' {
			result = append(result, i)
		}
	}

	return result
}
