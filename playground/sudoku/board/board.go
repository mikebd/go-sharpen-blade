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

// col is a vertical column of cells, the 0th cell is the topmost cell
type col [size]cell

// row is a horizontal row of cells, the 0th cell is the leftmost cell
type row [size]cell

// sct is a square of cells, the 0th cell is the top left cell.
// The 1th cell is the cell to the right of the 0th cell.
// The (size-1)th cell is the bottom right cell.
type sct [size]cell

// board is a grid of cells
type board [9]row
