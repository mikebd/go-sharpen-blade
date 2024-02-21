package cell

// CellCount is the number of cells in a row, column, or section.
// The square root of CellCount must be an integer.
// There must be a constant of type value for each value from 1 to CellCount.
const CellCount int = 9

// CellPointers is a slice of pointers to cells.  It allows navigation of
// the Board by multiple dimensions (e.g. row, column, section).
// An update to a cell in one dimension will be reflected in all dimensions.
type CellPointers []*Cell

type CellPointersArray [CellCount]CellPointers
