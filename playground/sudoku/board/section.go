package board

import (
	. "go-sharpen-blade/playground/sudoku/cell"
	"log"
	"math"
)

// sctSize is the number of rows and the number of columns in a section,
// i.e. the CellCount of a "side" of a section.
var sctSize = func() int {
	intSctSize := int(math.Sqrt(float64(CellCount)))
	if intSctSize*intSctSize != CellCount {
		log.Fatalf("CellCount must be a perfect square, but CellCount is %d", CellCount)
	}
	return intSctSize
}()

func (board *Board) section(row, col int) CellPointers {
	sectionOffset := col/sctSize + (row/sctSize)*sctSize
	return board.scts[sectionOffset]
}

func offsetWithinSection(row, col int) int {
	return col%sctSize + (row%sctSize)*sctSize
}
