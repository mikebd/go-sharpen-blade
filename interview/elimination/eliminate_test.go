package elimination

import (
	"reflect"
	"testing"
)

func Test_eliminateBoard(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]uint8
	}{
		{
			name: "135 246 789",
			args: args{
				input: []string{"135", "246", "789"},
			},
			want: [][]uint8{
				{1, 3, 5},
				{2, 4, 6},
				{7, 8, 9},
			},
		},
		{
			name: "111 234 567 891",
			args: args{
				input: []string{"111", "234", "567", "891"},
			},
			want: [][]uint8{
				{0, 0, 0},
				{2, 3, 4},
				{5, 6, 7},
				{8, 9, 1},
			},
		},
		{
			name: "1111 2345 5678 8912",
			args: args{
				input: []string{"1111", "2345", "5678", "8912"},
			},
			want: [][]uint8{
				{0, 0, 0, 0},
				{2, 3, 4, 5},
				{5, 6, 7, 8},
				{8, 9, 1, 2},
			},
		},
		{
			name: "2111 2345 5678 8912",
			args: args{
				input: []string{"2111", "2345", "5678", "8912"},
			},
			want: [][]uint8{
				{2, 0, 0, 0},
				{2, 3, 4, 5},
				{5, 6, 7, 8},
				{8, 9, 1, 2},
			},
		},
		{
			name: "111 234 267 291",
			args: args{
				input: []string{"111", "234", "267", "291"},
			},
			want: [][]uint8{
				{0, 0, 0},
				{0, 3, 4},
				{0, 6, 7},
				{0, 9, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminate(createBoard(tt.args.input)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
