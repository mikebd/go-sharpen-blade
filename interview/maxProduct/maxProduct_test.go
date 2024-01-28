package maxProduct

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1234", args{"1234"}, 24},
		{"12342", args{"12342"}, 48},
		{"12034210", args{"12034210"}, 24},
		{"12134211", args{"12134211"}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.s); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
