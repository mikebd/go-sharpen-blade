// https://leetcode.com/problems/count-and-say/

package countAndSay

import "strconv"

func countAndSay(n int) string {
	return say("1", n-1)
}

func say(s string, n int) string {
	if n == 0 {
		return s
	}

	result := ""
	count := 0
	prevDigit := ""
	for _, digitRune := range s {
		digit := string(digitRune)
		if digit == prevDigit {
			count++
		} else {
			if count > 0 {
				result += strconv.Itoa(count) + prevDigit
			}

			prevDigit = digit
			count = 1
		}
	}

	result += strconv.Itoa(count) + prevDigit

	return say(result, n-1)
}
