package day1

func Example_miniMaxSum_1() {
	miniMaxSum([]int32{1, 3, 5, 7, 9})

	// Output:
	// 16 24
}

func Example_miniMaxSum_2() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})

	// Output:
	// 10 14
}

func Example_miniMaxSum_3() {
	miniMaxSum([]int32{7, 69, 2, 221, 8974})

	// Output:
	// 299 9271
}
