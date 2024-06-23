package day2

func Example_diagonalDifference_1() {
	diagonalDifference([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	})

	// Output:
	// 2
}

func Example_diagonalDifference_3() {
	diagonalDifference([][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	})

	// Output:
	// 15
}
