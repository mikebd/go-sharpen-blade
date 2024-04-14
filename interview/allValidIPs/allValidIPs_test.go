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
		"1231231231234",
	}
	var want []string
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			if got := allValidIPs(tt); !reflect.DeepEqual(got, want) {
				t.Errorf("allValidIPs() = %v, want %v", got, want)
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
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := allValidIPs(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allValidIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
