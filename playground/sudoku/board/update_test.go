package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
)

func Example_board_setValue() {
	board := createEmptyBoard()

	_, err := board.setValue(0, 0, OneRuneValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(board.rows[0][0].Value()))
		fmt.Println(string(board.cols[0][0].Value()))
		fmt.Println(string(board.section(0, 0)[0].Value()))
	}

	// Output:
	// 1
	// 1
	// 1
}
