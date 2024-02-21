package cell

// External representation of Cell values, e.g. for creating
// or rendering a board.
const (
	NilRuneValue   rune = '?'
	EmptyRuneValue rune = ' '
	OneRuneValue   rune = '1'
	TwoRuneValue   rune = '2'
	ThreeRuneValue rune = '3'
	FourRuneValue  rune = '4'
	FiveRuneValue  rune = '5'
	SixRuneValue   rune = '6'
	SevenRuneValue rune = '7'
	EightRuneValue rune = '8'
	NineRuneValue  rune = '9'
)

func EmptyRune(runeValue rune) bool {
	return runeValue == NilRuneValue || runeValue == EmptyRuneValue
}

func runeToValue(runeValue rune) value {
	switch runeValue {
	case NilRuneValue:
		return nilValue
	case EmptyRuneValue:
		return empty
	case OneRuneValue:
		return one
	case TwoRuneValue:
		return two
	case ThreeRuneValue:
		return three
	case FourRuneValue:
		return four
	case FiveRuneValue:
		return five
	case SixRuneValue:
		return six
	case SevenRuneValue:
		return seven
	case EightRuneValue:
		return eight
	case NineRuneValue:
		return nine
	}
	return nilValue
}

func (value value) rune() rune {
	switch value {
	case nilValue:
		return NilRuneValue
	case empty:
		return EmptyRuneValue
	case one:
		return OneRuneValue
	case two:
		return TwoRuneValue
	case three:
		return ThreeRuneValue
	case four:
		return FourRuneValue
	case five:
		return FiveRuneValue
	case six:
		return SixRuneValue
	case seven:
		return SevenRuneValue
	case eight:
		return EightRuneValue
	case nine:
		return NineRuneValue
	}
	return NilRuneValue
}
