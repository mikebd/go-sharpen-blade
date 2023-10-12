package model

type Suit int8

const (
	UnknownSuit Suit = iota
	Heart
	Diamond
	Club
	Spade
)

func (s Suit) String() string {
	switch s {
	case Heart:
		return "♥"
	case Diamond:
		return "♦"
	case Club:
		return "♣"
	case Spade:
		return "♠"
	}
	return "?"
}
