package parseValidateIPs

import (
	"reflect"
	"testing"
)

func Test_extractPossibleIPs(t *testing.T) {
	type args struct {
		responseBody string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty response",
			args: args{""},
			want: []string{},
		},
		{
			name: "message with preamble",
			args: args{"{ \"timestamp\": 1710974865, \"message\": \"Here are the IPs: 123.12.32.1 and this also 10.100.12.3 256.32.23.1\" }"},
			want: []string{"123.12.32.1", "and", "this", "also", "10.100.12.3", "256.32.23.1"},
		},
		{
			name: "message without preamble",
			args: args{"{ \"timestamp\": 1710974865, \"message\": \"123.12.32.1 and this also 10.100.12.3 256.32.23.1\" }"},
			want: []string{"123.12.32.1", "and", "this", "also", "10.100.12.3", "256.32.23.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractPossibleIPs(tt.args.responseBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
