package canWin

import "testing"

type test struct {
	board          []int
	startingIndex  int
	expectedResult bool
}

func Test_canWin(t *testing.T) {
	tests := []test{
		{[]int{4, 1, 2, 0, 1}, 1, true},
		{[]int{2, 0, 3}, 0, false},
		{[]int{2, 1, 3}, 0, false},
		{[]int{4}, 0, false},
		{[]int{4}, 1, false},
		{[]int{1, 1, 1}, 1, false},
	}

	for _, tt := range tests {
		actualResult := canWin(tt.board, tt.startingIndex)
		if actualResult != tt.expectedResult {
			t.Errorf("canWin(%v, %v): expected %v, actual %v",
				tt.board, tt.startingIndex, tt.expectedResult, actualResult)
		}
	}
}
