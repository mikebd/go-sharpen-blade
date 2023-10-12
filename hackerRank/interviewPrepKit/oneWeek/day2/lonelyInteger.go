package day2

import "fmt"

// https://www.hackerrank.com/challenges/one-week-preparation-kit-lonely-integer

func lonelyInteger(arr []int32) {
	counts := make([]int8, 101)
	for _, n := range arr {
		counts[n]++
	}
	for i, count := range counts {
		if count == 1 {
			fmt.Println(i)
			break
		}
	}
}
