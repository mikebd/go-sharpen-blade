package model

type Deck interface {
	Cards() *Cards
	Shuffle()
}
