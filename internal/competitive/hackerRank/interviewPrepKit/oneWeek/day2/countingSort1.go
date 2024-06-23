package day2

// https://www.hackerrank.com/challenges/one-week-preparation-kit-countingsort1

const resultSize = 100

func countingSort1(arr []int32) []int32 {
	result := make([]int32, resultSize)

	for _, val := range arr {
		result[val]++
	}

	return result
}
