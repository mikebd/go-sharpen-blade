package board

import (
	"fmt"
	"testing"
)

func Example_board_validate_valid() {
	board := create([size]string{})

	err := board.setValue(0, 0, one)
	if err != nil {
		fmt.Println(err)
	}

	invalid := board.validate()

	fmt.Println(invalid.any())

	// Output:
	// false
}

func Test_board_validate_invalid(t *testing.T) {
	type setValueArgs struct {
		row, col int
		value    value
	}

	type args struct {
		setValueArgs []setValueArgs
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"invalid row - duplicate value",
			args{
				[]setValueArgs{
					{0, 0, one},
					{0, 1, one},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := create([size]string{})

			for _, setValueArgs := range tt.args.setValueArgs {
				err := board.setValue(setValueArgs.row, setValueArgs.col, setValueArgs.value)
				if err != nil {
					t.Error(err)
				}
			}

			invalid := board.validate()

			if !invalid.any() {
				t.Error("expected invalid.any() to be true")
			}
		})
	}
}
