package board

import "fmt"

func Example_board_setValue() {
	board := create([size]string{})

	err := board.setValue(0, 0, one)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(board.rows[0][0].value)
		fmt.Println(board.cols[0][0].value)
		fmt.Println(board.section(0, 0)[0].value)
	}

	// Output:
	// 1
	// 1
	// 1
}
