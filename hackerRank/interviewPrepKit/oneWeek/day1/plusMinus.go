package day1

import "fmt"

// https://www.hackerrank.com/challenges/one-week-preparation-kit-plus-minus

func plusMinus(arr []int32) {
	countPositive := 0
	countNegative := 0
	countZero := 0

	for _, num := range arr {
		if num > 0 {
			countPositive++
		} else if num < 0 {
			countNegative++
		} else {
			countZero++
		}
	}

	countTotal := float64(len(arr))
	ratioPositive := float64(countPositive) / countTotal
	ratioNegative := float64(countNegative) / countTotal
	ratioZero := float64(countZero) / countTotal
	fmt.Printf("%.6f\n", ratioPositive)
	fmt.Printf("%.6f\n", ratioNegative)
	fmt.Printf("%.6f\n", ratioZero)
}
