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

	type args struct {
		cells       cellPointersArray
		resultField []int
	}

	lop.ForEach([]args{
		{board.rows, result.rows},
		{board.cols, result.cols},
		{board.scts, result.scts},
	},
		func(args args, _ int) {
			lop.ForEach([]cellPointersArray{args.cells}, func(cells cellPointersArray, col int) {
				if invalidCells(cells) {
					args.resultField = append(args.resultField, col)
				}
			})
		})

	return result
}

func invalidCells(cells cellPointersArray) bool {
	return true
}
