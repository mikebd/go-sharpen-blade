package day1

func Example_timeConversion_1() {
	timeConversion("12:01:00PM")

	// Output:
	// 12:01:00
}

func Example_timeConversion_2() {
	timeConversion("12:01:00AM")

	// Output:
	// 00:01:00
}

func Example_timeConversion_3() {
	timeConversion("07:05:45PM")

	// Output:
	// 19:05:45
}
