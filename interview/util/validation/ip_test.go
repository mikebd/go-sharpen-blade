package validation

import "testing"

func TestIsValidIPv4(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{""}, false},
		{args{"123"}, false},
		{args{"123.123"}, false},
		{args{"123.123.123"}, false},
		{args{"123.123.123.123.123"}, false},

		{args{"123.123.123.123"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.args.ip, func(t *testing.T) {
			if got := IsValidIPv4(tt.args.ip); got != tt.want {
				t.Errorf("IsValidIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidIPv4Segment(t *testing.T) {
	type args struct {
		segment string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{""}, false},
		{args{"00"}, false},
		{args{"001"}, false},
		{args{"a1"}, false},
		{args{"1a"}, false},
		{args{"1234"}, false},
		{args{"256"}, false},
		{args{"-1"}, false},

		{args{"0"}, true},
		{args{"1"}, true},
		{args{"255"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.args.segment, func(t *testing.T) {
			if got := IsValidIPv4Segment(tt.args.segment); got != tt.want {
				t.Errorf("IsValidIPv4Segment() = %v, want %v", got, tt.want)
			}
		})
	}
}
