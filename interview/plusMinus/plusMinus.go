package plusMinus

func Solution(S string) int {
	result := 0
	operator := '+'

	for i := 0; i < len(S); i++ {
		switch {
		case S[i] == 'o' || S[i] == 'O':
			if operator == '+' {
				result += 1
			} else {
				result -= 1
			}
			i += 2
		case S[i] == 't' || S[i] == 'T':
			if operator == '+' {
				result += 2
			} else {
				result -= 2
			}
			i += 2
		case S[i] == '+':
			operator = '+'
		case S[i] == '-':
			operator = '-'
		}
	}

	return result
}
