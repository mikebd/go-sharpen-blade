package board

// value is a compressed storage type for a digit in a cell.
type value uint8

// Internal representation of each value that can be stored in a cell.
const (
	empty value = iota
	one
	two
	three
	four
	five
	six
	seven
	eight
	nine
	// Add more values here to support larger boards
	nilValue value = 255 // used to indicate an invalid value
)

// External representation of cell values, e.g. for creating
// or rendering a board.
const (
	nilRuneValue   rune = '?'
	emptyRuneValue rune = ' '
	oneRuneValue   rune = '1'
	twoRuneValue   rune = '2'
	threeRuneValue rune = '3'
	fourRuneValue  rune = '4'
	fiveRuneValue  rune = '5'
	sixRuneValue   rune = '6'
	sevenRuneValue rune = '7'
	eightRuneValue rune = '8'
	nineRuneValue  rune = '9'
)

func runeToValue(runeValue rune) value {
	switch runeValue {
	case nilRuneValue:
		return nilValue
	case emptyRuneValue:
		return empty
	case oneRuneValue:
		return one
	case twoRuneValue:
		return two
	case threeRuneValue:
		return three
	case fourRuneValue:
		return four
	case fiveRuneValue:
		return five
	case sixRuneValue:
		return six
	case sevenRuneValue:
		return seven
	case eightRuneValue:
		return eight
	case nineRuneValue:
		return nine
	}
	return nilValue
}

func (value value) rune() rune {
	switch value {
	case nilValue:
		return nilRuneValue
	case empty:
		return emptyRuneValue
	case one:
		return oneRuneValue
	case two:
		return twoRuneValue
	case three:
		return threeRuneValue
	case four:
		return fourRuneValue
	case five:
		return fiveRuneValue
	case six:
		return sixRuneValue
	case seven:
		return sevenRuneValue
	case eight:
		return eightRuneValue
	case nine:
		return nineRuneValue
	}
	return nilRuneValue
}
