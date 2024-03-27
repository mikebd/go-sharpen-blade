package fibonacciNumber

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 5},
		{"6", args{6}, 8},
		{"7", args{7}, 13},
		{"8", args{8}, 21},
		{"9", args{9}, 34},
		{"10", args{10}, 55},
		{"11", args{11}, 89},
		{"12", args{12}, 144},
		{"13", args{13}, 233},
		{"14", args{14}, 377},
		{"15", args{15}, 610},
		{"16", args{16}, 987},
		{"17", args{17}, 1597},
		{"18", args{18}, 2584},
		{"19", args{19}, 4181},
		{"20", args{20}, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
