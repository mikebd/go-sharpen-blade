package cell

// value is a compressed storage type for a digit in a Cell.
type value uint8

// Internal representation of each value that can be stored in a Cell.
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
