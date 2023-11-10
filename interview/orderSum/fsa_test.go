package orderSum

import "fmt"

func Example_sumByFsaByQuarter() {
	orders := []order{
		{
			"2022-01-01", "M5C 2N4", "Joe Smith", 100,
		},
	}

	result := sumByFsaByQuarter(orders)

	fmt.Println(result)

	// Output:
	// map[Q1:map[M5C:100] Q2:map[] Q3:map[] Q4:map[]]
}

func Example_sumByFsa1() {
	orders := []order{
		{
			"2022-01-01", "M5C 2N4", "Joe Smith", 100,
		},
	}

	result := sumByFsa(orders)

	fmt.Println(result)

	// Output:
	// map[M5C:100]
}

func Example_sumByFsa() {
	orders := []order{
		{
			"2022-01-01", "M5C 2N4", "Joe Smith", 100,
		},
	}

	result := sumByFsa(orders)

	fmt.Println(result)

	// Output:
	// map[M5C:100]
}

/*
[Timestamp, Postal Code, Name, Volume]
[
["2022-01-01", "M5C 2N4", "Joe Smith", 100],
["2022-01-01", "M5C 9V1", "Jane Doe", 50],
["2022-01-01", "L7Q 5J1", "Barbara Shoe", 25],
["2022-02-01", "M5C 2N4", "Joe Smith", 101],
["2022-02-01", "M5C 9V1", "Jane Doe", 51],
["2022-02-01", "L7Q 5J1", "Barbara Shoe", 26],
["2022-03-01", "M5C 2N4", "Joe Smith", 102],
["2022-03-01", "M5C 9V1", "Jane Doe", 52],
["2022-03-01", "L7Q 5J1", "Barbara Shoe", 27],
["2022-04-01", "M5C 2N4", "Joe Smith", 103],
["2022-04-01", "M5C 9V1", "Jane Doe", 53],
["2022-04-01", "L7Q 5J1", "Barbara Shoe", 28],
]

Return the sums of volume aggregated by FSA.
Note: The "FSA" is the first three characters of a postal code

E.g.

{
"M5C": 456,
"L7Q": 78,
}

*/
