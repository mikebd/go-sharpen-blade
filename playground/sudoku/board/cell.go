package board

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
