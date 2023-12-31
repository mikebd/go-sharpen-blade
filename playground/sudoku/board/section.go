package board

import (
	"log"
	"math"
)

// sctSize is the number of rows and the number of columns in a section,
// i.e. the size of a "side" of a section.
var sctSize = func() int {
	intSctSize := int(math.Sqrt(float64(size)))
	if intSctSize*intSctSize != size {
		log.Fatalf("size must be a perfect square, but size is %d", size)
	}
	return intSctSize
}()

func (board *Board) section(row, col int) cellPointers {
	sectionOffset := col/sctSize + (row/sctSize)*sctSize
	return board.scts[sectionOffset]
}

func offsetWithinSection(row, col int) int {
	return col%sctSize + (row%sctSize)*sctSize
}
