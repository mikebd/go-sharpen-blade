package board

import (
	lop "github.com/samber/lo/parallel"
)

type invalid struct {
	rows, cols, scts []int
}

func (invalid *invalid) any() bool {
	return invalid.rows != nil || invalid.cols != nil || invalid.scts != nil
}

func (board *board) validate() invalid {
	result := invalid{}

	lop.ForEach([]cellPointersArray{board.rows}, func(cells cellPointersArray, row int) {
		if invalidCells(cells) {
			result.rows = append(result.rows, row)
		}
	})

	lop.ForEach([]cellPointersArray{board.cols}, func(cells cellPointersArray, col int) {
		if invalidCells(cells) {
			result.cols = append(result.cols, col)
		}
	})

	lop.ForEach([]cellPointersArray{board.scts}, func(cells cellPointersArray, sct int) {
		if invalidCells(cells) {
			result.scts = append(result.scts, sct)
		}
	})

	return result
}

func invalidCells(cells cellPointersArray) bool {
	return true
}
