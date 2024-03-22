package plusMinus

import "log"

func plusMinus(expr string) int {
	result := 0
	operator := '+'

	for i := 0; i < len(expr); i++ {
		switch expr[i] {
		case 'o', 'O':
			result = operate(result, operator, 1)
			i += 2
		case 'e', 'E':
			result = operate(result, operator, 8)
			i += 4
		case 'f', 'F':
			switch expr[i+1] {
			case 'i', 'I':
				result = operate(result, operator, 5)
				i += 3
			case 'o', 'O':
				result = operate(result, operator, 4)
				i += 3
			}
		case 'n', 'N':
			result = operate(result, operator, 9)
			i += 3
		case 's', 'S':
			switch expr[i+1] {
			case 'i', 'I':
				result = operate(result, operator, 6)
				i += 2
			case 'e', 'E':
				result = operate(result, operator, 7)
				i += 4
			}
		case 't', 'T':
			switch expr[i+1] {
			case 'w', 'W':
				result = operate(result, operator, 2)
				i += 2
			case 'h', 'H':
				result = operate(result, operator, 3)
				i += 4
			}
		case 'z', 'Z':
			result = operate(result, operator, 0)
			i += 3
		case '+':
			operator = '+'
		case '-':
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
