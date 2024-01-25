package maxProduct

/*
	Given a string of decimal digits, output an integer of the largest
	product of any four consecutive digits.
*/

const (
	consecutiveDigits = 4
)

func maxProduct(s string) int {
	return nestedSolution(s)
}

// nestedSolution is the least efficient solution.  I used a sliding window
// approach in the interview but have captured this here for benchmarking.
func nestedSolution(s string) int {
	maxProduct := 0
	for i := 0; i <= len(s)-consecutiveDigits; i++ {
		product := 1
		for j := i; j < i+consecutiveDigits; j++ {
			product *= stringRuneToDecimal(s, j)
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct
}

// stringRuneToInt converts a string rune to a decimal integer or 0 if invalid
func stringRuneToDecimal(s string, i int) int {
	if i < 0 || i >= len(s) {
		return 0
	}

	val := int(s[i] - '0')

	if val < 0 || val > 9 {
		return 0
	}

	return val
}
