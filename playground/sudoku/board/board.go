package board

// size is the number of cells in a row, column, or section.
// The square root of size must be an integer.
// There must be a constant of type value for each value from 1 to size.
const size int = 9

type cellPointersArray [size]cellPointers

// board has cell accessors
type board struct {
	// rows is an array of horizontal rows of cells, the 0th cell is the leftmost cell.
	rows cellPointersArray

	// cols is an array of vertical columns of cells, the 0th cell is the topmost cell.
	cols cellPointersArray

	// scts is an array of squares of cells, the 0th cell is the top left cell.
	// The 1th cell is the cell to the right of the 0th cell.
	// The (size-1)th cell is the bottom right cell.
	scts cellPointersArray
}
