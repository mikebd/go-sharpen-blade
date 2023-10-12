package day1

func Example_lonelyInteger_1() {
	lonelyInteger([]int32{1, 2, 2, 1, 4})

	// Output:
	// 4
}

func Example_lonelyInteger_2() {
	lonelyInteger([]int32{0, 0, 1, 3, 2, 40, 2, 1, 3})

	// Output:
	// 40
}

func Example_lonelyInteger_3() {
	lonelyInteger([]int32{8, 3, 5, 1, 5, 3, 8, 100, 100})

	// Output:
	// 1
}
