package model

type Rank int8

const (
	UnknownRank Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (r Rank) String() string {
	switch r {
	case Ace:
		return " A"
	case Two:
		return " 2"
	case Three:
		return " 3"
	case Four:
		return " 4"
	case Five:
		return " 5"
	case Six:
		return " 6"
	case Seven:
		return " 7"
	case Eight:
		return " 8"
	case Nine:
		return " 9"
	case Ten:
		return "10"
	case Jack:
		return " J"
	case Queen:
		return " Q"
	case King:
		return " K"
	}
	return " ?"
}
