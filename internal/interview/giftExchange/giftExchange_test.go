package giftExchange

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGiftExchange(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		wantLen    int
		errRequire require.ErrorAssertionFunc
	}{
		{
			name:       "valid input - 3 people",
			filename:   "valid_3",
			wantLen:    3,
			errRequire: require.NoError,
		},
		{
			name:       "invalid input - 1 person",
			filename:   "invalid_1",
			wantLen:    0,
			errRequire: require.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := testDataFile(t, tt.filename)

			got, err := giftExchange(f)

			tt.errRequire(t, err)
			require.Equal(t, tt.wantLen, len(got), "expected gift exchange list length to be %d but got %d", tt.wantLen, len(got))

			for i, sr := range got {
				assert.NotEqual(t, sr.sender, sr.receiver, "sender and receiver should be different for element %d", i)
			}

			// Verify each sender is unique and each receiver is unique
			senders := make(map[person]bool)
			receivers := make(map[person]bool)
			for i, sr := range got {
				assert.False(t, senders[sr.sender], "sender should be unique for element %d", i)
				assert.False(t, receivers[sr.receiver], "receiver should be unique for element %d", i)
				senders[sr.sender] = true
				receivers[sr.receiver] = true
			}
		})
	}
}

func testDataFile(t *testing.T, name string) *os.File {
	t.Helper()

	path := filepath.Join("testdata", name+".csv")
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	return f
}
