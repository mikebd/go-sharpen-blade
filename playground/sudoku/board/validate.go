package board

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go-sharpen-blade/playground/sudoku/cell"
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
		cells       cell.CellPointersArray
		resultField *[]int
	}

	lop.ForEach([]args{
		{board.rows, &result.rows},
		{board.cols, &result.cols},
		{board.scts, &result.scts},
	},
		func(args args, _ int) {
			lop.ForEach([]cell.CellPointersArray{args.cells}, func(cells cell.CellPointersArray, index int) {
				if invalidCells(cells[index]) {
					*args.resultField = append(*args.resultField, index)
				}
			})
		})

	return &result
}

func invalidCells(cells cell.CellPointers) bool {
	// duplicate values
	// count the number of times each value appears
	// if any value appears more than once, the cells are invalid
	counts := lo.CountValuesBy(cells, cell.Value)

	for runeValue, count := range counts {
		if !cell.EmptyRune(runeValue) {
			if count > 1 {
				return true
			}
		}
	}

	return false
}
