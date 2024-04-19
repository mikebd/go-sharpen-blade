package canBeMadeValidParentheses

import (
	"reflect"
	"testing"
)

func Test_canBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		args args
	}{
		{args{"((", "00"}},
		{args{"))", "00"}},
		{args{"))((", "0000"}},
		{args{"))()))", "010100"}},
		// fails: {args{"))()))))", "01010001"}},
		{args{"()()", "0000"}},
	}
	for _, tt := range tests {
		t.Run(tt.args.s+" "+tt.args.locked, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); !got {
				t.Errorf("canBeValid() = %v, want true", got)
			}
		})
	}
}

func Test_canNotBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		args args
	}{
		{args{")", "0"}},
		{args{"))", "10"}},
	}
	for _, tt := range tests {
		t.Run(tt.args.s+" "+tt.args.locked, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); got {
				t.Errorf("canBeValid() = %v, want false", got)
			}
		})
	}
}

func Test_unlockedOffsets(t *testing.T) {
	type args struct {
		locked string
	}
	tests := []struct {
		args args
		want []int
	}{
		{args{""}, []int{}},
		{args{"1"}, []int{}},
		{args{"0"}, []int{0}},
		{args{"10"}, []int{1}},
		{args{"010"}, []int{0, 2}},
		{args{"101"}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.args.locked, func(t *testing.T) {
			if got := unlockedOffsets(tt.args.locked); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unlockedOffsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
