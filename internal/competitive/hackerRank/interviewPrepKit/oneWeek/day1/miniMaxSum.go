package day1

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum

func miniMaxSum(arr []int32) {
	minimum := int32(math.MaxInt32)
	maximum := int32(math.MinInt32)
	var sum int64

	for _, num := range arr {
		minimum = min(num, minimum)
		maximum = max(num, maximum)
		sum += int64(num)
	}

	minSum := sum - int64(maximum)
	maxSum := sum - int64(minimum)
	fmt.Println(minSum, maxSum)
}
