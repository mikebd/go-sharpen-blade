package board

type value uint8

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
)

type cell struct {
	row, col     int
	value        value
	initialValue value
}

// size is the number of cells in a row, column, or section.
// The square root of size must be an integer.
// There must be a constant of type value for each value from 1 to size.
const size int = 9

// row is a horizontal row of cells, the 0th cell is the leftmost cell.
// row has the actual cells, not pointers to cells.
type row [size]cell

// col is a vertical column of cells, the 0th cell is the topmost cell.
// col has pointers to cells, not the actual cells.
type col [size]*cell

// sct is a square of cells, the 0th cell is the top left cell.
// The 1th cell is the cell to the right of the 0th cell.
// The (size-1)th cell is the bottom right cell.
// sct has pointers to cells, not the actual cells.
type sct [size]*cell

// board has cell accessors
type board struct {
	rows [size]row
	cols [size]col
	// scts [size]sct
}
