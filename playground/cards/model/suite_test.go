package model

import "fmt"

func Example_suiteToString() {
	fmt.Println(UnknownSuit)
	fmt.Println(Heart)
	fmt.Println(Diamond)
	fmt.Println(Club)
	fmt.Println(Spade)

	// Output:
	// ?
	// ♥
	// ♦
	// ♣
	// ♠
}
