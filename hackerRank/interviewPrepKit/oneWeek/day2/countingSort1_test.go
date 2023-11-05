package day2

import "testing"

func testIndexValue(t *testing.T, arr []int32, index int, expected int32) {
	if arr[index] != expected {
		t.Errorf("countingSort1() [%v] = %v, want %v", index, arr[index], expected)
	}
}

func Test_countingSort1(t *testing.T) {
	input := []int32{1, 1, 3, 2, 1}
	result := countingSort1(input)

	if len(result) != resultSize {
		t.Errorf("countingSort1() len(arr) = %v, want %v", len(result), resultSize)
	}

	testIndexValue(t, result, 0, 0)
	testIndexValue(t, result, 1, 3)
	testIndexValue(t, result, 2, 1)
	testIndexValue(t, result, 3, 1)

	for i := 4; i < 100; i++ {
		testIndexValue(t, result, i, 0)
	}
}
