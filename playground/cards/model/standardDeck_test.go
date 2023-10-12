package model

import "fmt"

func Example_standardDeck() {
	deck := NewStandardDeck()
	cards := *deck.Cards()
	fmt.Println(len(cards))

	// Output:
	// 52
}
