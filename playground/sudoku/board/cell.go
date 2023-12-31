package board

// value is a compressed storage type for a digit in a cell.
type value uint8

// Each digit that can be stored in a cell.
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

// cell is a single cell in a Board.
type cell struct {
	row, col     int
	value        value
	initialValue value
}

// cellPointers is a slice of pointers to cells.  It allows navigation of
// the Board by multiple dimensions (e.g. row, column, section).
// An update to a cell in one dimension will be reflected in all of the
// other dimensions.
type cellPointers []*cell

// empty returns true if the cell is empty.
func (cell *cell) empty() bool {
	return cell.value == empty
}

// setValue sets the value of the cell and returns the previous value.
func (cell *cell) setValue(value value) value {
	previousValue := cell.value
	cell.value = value
	return previousValue
}
