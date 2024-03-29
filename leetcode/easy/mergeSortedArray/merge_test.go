package mergeSortedArray

import (
	"slices"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1",
			args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			[]int{1, 2, 2, 3, 5, 6}},
		{"Example 2",
			args{[]int{1}, 1, []int{}, 0},
			[]int{1}},
		{"Example 3",
			args{[]int{0}, 0, []int{1}, 1},
			[]int{1}},
		{"12/59 tests pass",
			args{[]int{2, 0}, 1, []int{1}, 1},
			[]int{1, 2}},
		{"16/59 tests pass",
			args{[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3},
			[]int{-1, 0, 0, 1, 2, 2, 3, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !slices.Equal(tt.args.nums1, tt.want) {
				t.Errorf("got %v, wanted %v", tt.args.nums1, tt.want)
			}
		})
	}
}
