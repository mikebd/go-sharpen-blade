package board

import (
	"reflect"
	"testing"
)

func Test_create(t *testing.T) {
	type args struct {
		initialValues [size]string
	}
	tests := []struct {
		name string
		args args
		want board
	}{
		{
			name: "empty board",
			args: args{[size]string{}},
			want: board{
				row{
					cell{0, 0, empty, empty},
					cell{0, 1, empty, empty},
					cell{0, 2, empty, empty},
					cell{0, 3, empty, empty},
					cell{0, 4, empty, empty},
					cell{0, 5, empty, empty},
					cell{0, 6, empty, empty},
					cell{0, 7, empty, empty},
					cell{0, 8, empty, empty}},
				row{
					cell{1, 0, empty, empty},
					cell{1, 1, empty, empty},
					cell{1, 2, empty, empty},
					cell{1, 3, empty, empty},
					cell{1, 4, empty, empty},
					cell{1, 5, empty, empty},
					cell{1, 6, empty, empty},
					cell{1, 7, empty, empty},
					cell{1, 8, empty, empty}},
				row{
					cell{2, 0, empty, empty},
					cell{2, 1, empty, empty},
					cell{2, 2, empty, empty},
					cell{2, 3, empty, empty},
					cell{2, 4, empty, empty},
					cell{2, 5, empty, empty},
					cell{2, 6, empty, empty},
					cell{2, 7, empty, empty},
					cell{2, 8, empty, empty}},
				row{
					cell{3, 0, empty, empty},
					cell{3, 1, empty, empty},
					cell{3, 2, empty, empty},
					cell{3, 3, empty, empty},
					cell{3, 4, empty, empty},
					cell{3, 5, empty, empty},
					cell{3, 6, empty, empty},
					cell{3, 7, empty, empty},
					cell{3, 8, empty, empty}},
				row{
					cell{4, 0, empty, empty},
					cell{4, 1, empty, empty},
					cell{4, 2, empty, empty},
					cell{4, 3, empty, empty},
					cell{4, 4, empty, empty},
					cell{4, 5, empty, empty},
					cell{4, 6, empty, empty},
					cell{4, 7, empty, empty},
					cell{4, 8, empty, empty}},
				row{
					cell{5, 0, empty, empty},
					cell{5, 1, empty, empty},
					cell{5, 2, empty, empty},
					cell{5, 3, empty, empty},
					cell{5, 4, empty, empty},
					cell{5, 5, empty, empty},
					cell{5, 6, empty, empty},
					cell{5, 7, empty, empty},
					cell{5, 8, empty, empty}},
				row{
					cell{6, 0, empty, empty},
					cell{6, 1, empty, empty},
					cell{6, 2, empty, empty},
					cell{6, 3, empty, empty},
					cell{6, 4, empty, empty},
					cell{6, 5, empty, empty},
					cell{6, 6, empty, empty},
					cell{6, 7, empty, empty},
					cell{6, 8, empty, empty}},
				row{
					cell{7, 0, empty, empty},
					cell{7, 1, empty, empty},
					cell{7, 2, empty, empty},
					cell{7, 3, empty, empty},
					cell{7, 4, empty, empty},
					cell{7, 5, empty, empty},
					cell{7, 6, empty, empty},
					cell{7, 7, empty, empty},
					cell{7, 8, empty, empty}},
				row{
					cell{8, 0, empty, empty},
					cell{8, 1, empty, empty},
					cell{8, 2, empty, empty},
					cell{8, 3, empty, empty},
					cell{8, 4, empty, empty},
					cell{8, 5, empty, empty},
					cell{8, 6, empty, empty},
					cell{8, 7, empty, empty},
					cell{8, 8, empty, empty},
				},
			},
		}, {
			name: "three of each value",
			args: args{[size]string{
				"1  2  3",
				"4  5  6",
				"7  8  9",
				" 1  2  3",
				" 4  5  6",
				" 7  8  9",
				"  1  2  3",
				"  4  5  6",
				"  7  8  9",
			}},
			want: board{
				row{
					cell{0, 0, one, one},
					cell{0, 1, empty, empty},
					cell{0, 2, empty, empty},
					cell{0, 3, two, two},
					cell{0, 4, empty, empty},
					cell{0, 5, empty, empty},
					cell{0, 6, three, three},
					cell{0, 7, empty, empty},
					cell{0, 8, empty, empty}},
				row{
					cell{1, 0, four, four},
					cell{1, 1, empty, empty},
					cell{1, 2, empty, empty},
					cell{1, 3, five, five},
					cell{1, 4, empty, empty},
					cell{1, 5, empty, empty},
					cell{1, 6, six, six},
					cell{1, 7, empty, empty},
					cell{1, 8, empty, empty}},
				row{
					cell{2, 0, seven, seven},
					cell{2, 1, empty, empty},
					cell{2, 2, empty, empty},
					cell{2, 3, eight, eight},
					cell{2, 4, empty, empty},
					cell{2, 5, empty, empty},
					cell{2, 6, nine, nine},
					cell{2, 7, empty, empty},
					cell{2, 8, empty, empty}},
				row{
					cell{3, 0, empty, empty},
					cell{3, 1, one, one},
					cell{3, 2, empty, empty},
					cell{3, 3, empty, empty},
					cell{3, 4, two, two},
					cell{3, 5, empty, empty},
					cell{3, 6, empty, empty},
					cell{3, 7, three, three},
					cell{3, 8, empty, empty}},
				row{
					cell{4, 0, empty, empty},
					cell{4, 1, four, four},
					cell{4, 2, empty, empty},
					cell{4, 3, empty, empty},
					cell{4, 4, five, five},
					cell{4, 5, empty, empty},
					cell{4, 6, empty, empty},
					cell{4, 7, six, six},
					cell{4, 8, empty, empty}},
				row{
					cell{5, 0, empty, empty},
					cell{5, 1, seven, seven},
					cell{5, 2, empty, empty},
					cell{5, 3, empty, empty},
					cell{5, 4, eight, eight},
					cell{5, 5, empty, empty},
					cell{5, 6, empty, empty},
					cell{5, 7, nine, nine},
					cell{5, 8, empty, empty}},
				row{
					cell{6, 0, empty, empty},
					cell{6, 1, empty, empty},
					cell{6, 2, one, one},
					cell{6, 3, empty, empty},
					cell{6, 4, empty, empty},
					cell{6, 5, two, two},
					cell{6, 6, empty, empty},
					cell{6, 7, empty, empty},
					cell{6, 8, three, three}},
				row{
					cell{7, 0, empty, empty},
					cell{7, 1, empty, empty},
					cell{7, 2, four, four},
					cell{7, 3, empty, empty},
					cell{7, 4, empty, empty},
					cell{7, 5, five, five},
					cell{7, 6, empty, empty},
					cell{7, 7, empty, empty},
					cell{7, 8, six, six}},
				row{
					cell{8, 0, empty, empty},
					cell{8, 1, empty, empty},
					cell{8, 2, seven, seven},
					cell{8, 3, empty, empty},
					cell{8, 4, empty, empty},
					cell{8, 5, eight, eight},
					cell{8, 6, empty, empty},
					cell{8, 7, empty, empty},
					cell{8, 8, nine, nine},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := create(tt.args.initialValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("create() = %v, want %v", got, tt.want)
			}
		})
	}
}
