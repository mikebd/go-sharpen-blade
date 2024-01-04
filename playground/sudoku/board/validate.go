package board

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

type Invalid struct {
	rows, cols, scts []int
}

func (invalid *Invalid) Any() bool {
	return invalid.rows != nil || invalid.cols != nil || invalid.scts != nil
}

func (board *Board) Validate() *Invalid {
	result := Invalid{}

	type args struct {
		cells       cellPointersArray
		resultField *[]int
	}

	lop.ForEach([]args{
		{board.rows, &result.rows},
		{board.cols, &result.cols},
		{board.scts, &result.scts},
	},
		func(args args, _ int) {
			lop.ForEach([]cellPointersArray{args.cells}, func(cells cellPointersArray, index int) {
				if invalidCells(cells[index]) {
					*args.resultField = append(*args.resultField, index)
				}
			})
		})

	return &result
}

func invalidCells(cells cellPointers) bool {
	cellValue := func(cell *cell) value {
		if cell == nil {
			return nilValue
		}
		return cell.value
	}

	// duplicate values
	// count the number of times each value appears
	// if any value appears more than once, the cells are invalid
	counts := lo.CountValuesBy(cells, cellValue)

	for value, count := range counts {
		if value != empty && count > 1 {
			return true
		}
	}

	return false
}
