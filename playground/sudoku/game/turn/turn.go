package turn

// Turn captures a single move made in the game
type Turn interface {
	Number() int

	Error() error

	Valid() bool

	Row() int

	Col() int

	PreviousValue() rune

	Value() rune
}
