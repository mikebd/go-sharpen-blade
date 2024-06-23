package day2

import "fmt"

// https://www.hackerrank.com/challenges/one-week-preparation-kit-diagonal-difference

func diagonalDifference(arr [][]int32) {
	sum1, sum2 := int32(0), int32(0)
	for i := range arr[0] {
		sum1 += arr[i][i]
		sum2 += arr[len(arr)-i-1][i]
	}
	if sum1 > sum2 {
		fmt.Println(sum1 - sum2)
	} else {
		fmt.Println(sum2 - sum1)
	}
}
