package port

// Turn captures a single move made in the game
type Turn interface {
	Number() int

	Valid() bool

	Row() int

	Col() int

	PreviousValue() rune

	Value() rune
}
