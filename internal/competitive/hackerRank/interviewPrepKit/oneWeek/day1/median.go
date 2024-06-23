package day1

import (
	"fmt"
	"sort"
)

func median(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	isOdd := len(arr)%2 == 1
	if isOdd {
		fmt.Println(arr[len(arr)/2])
	} else {
		fmt.Println((arr[len(arr)/2-1] + arr[len(arr)/2]) / 2)
	}
}
