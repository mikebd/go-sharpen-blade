package day1

func Example_plusMinus_1() {
	plusMinus([]int32{1, 1, 0, -1, -1})

	// Output:
	// 0.400000
	// 0.400000
	// 0.200000
}

func Example_plusMinus_2() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})

	// Output:
	// 0.500000
	// 0.333333
	// 0.166667
}

func Example_plusMinus_3() {
	plusMinus([]int32{1, 2, 3, -1, -2, -3, 0, 0})

	// Output:
	// 0.375000
	// 0.375000
	// 0.250000
}
