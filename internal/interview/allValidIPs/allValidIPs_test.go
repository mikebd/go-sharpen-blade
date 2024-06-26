package allValidIPs

import (
	"reflect"
	"testing"
)

func Test_allValidIPsInvalid(t *testing.T) {
	tests := []string{
		"",
		"1",
		"12",
		"123",
		"1a23",
		"00000",
		"1231231231234",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			if got := allValidIPs(tt); len(got) != 0 {
				t.Errorf("allValidIPs() = %v", got)
			}
		})
	}
}

func Test_allValidIPsValid(t *testing.T) {
	tests := []struct {
		arg  string
		want []string
	}{
		{"0000", []string{"0.0.0.0"}},
		{"1111", []string{"1.1.1.1"}},
		{"11111", []string{
			"1.1.1.11",
			"1.1.11.1",
			"1.11.1.1",
			"11.1.1.1",
		}},
		{"123123123123", []string{
			"123.123.123.123",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := allValidIPs(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allValidIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
