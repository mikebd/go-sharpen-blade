package plusMinus

import "log"

func plusMinus(expr string) int {
	result := 0
	operator := '+'

	for i := 0; i < len(expr); i++ {
		switch {
		case expr[i] == 'o' || expr[i] == 'O':
			result = operate(result, operator, 1)
			i += 2
		case expr[i] == 't' || expr[i] == 'T':
			result = operate(result, operator, 2)
			i += 2
		case expr[i] == '+':
			operator = '+'
		case expr[i] == '-':
			operator = '-'
		}
	}

	return result
}

func operate(currentValue int, operator rune, operand int) int {
	result := currentValue

	switch operator {
	case '+':
		result += operand
	case '-':
		result -= operand
	default:
		log.Fatalln("Invalid operator", operator)
	}

	return result
}
