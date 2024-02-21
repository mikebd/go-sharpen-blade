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
		fmt.Println(board.rows[0][0].Value() - '0')
		fmt.Println(board.cols[0][0].Value() - '0')
		fmt.Println(board.section(0, 0)[0].Value() - '0')
	}

	// Output:
	// 1
	// 1
	// 1
}
