package board

import (
	"fmt"
	. "go-sharpen-blade/playground/sudoku/cell"
	"testing"
)

func Example_board_validate_valid() {
	board := createEmptyBoard()

	_, err := board.setValue(0, 0, OneRuneValue)
	if err != nil {
		fmt.Println(err)
	}

	invalid := board.Validate()

	fmt.Println(invalid.Any())

	// Output:
	// false
}

var correctCompletedBoard = func() *board {
	board, err := create([CellCount]string{
		"534678912",
		"672195348",
		"198342567",
		"859761423",
		"426853791",
		"713924856",
		"961537284",
		"287419635",
		"345286179",
	})

	if err != nil {
		panic(err)
	}

	return board
}()

func Example_board_validate_correctCompletedBoard() {
	invalid := correctCompletedBoard.Validate()

	fmt.Println(invalid.Any())

	// Output:
	// false
}

func Test_board_validate_invalid(t *testing.T) {
	type setValueArgs struct {
		row, col int
		value    rune
	}

	type args struct {
		setValueArgs []setValueArgs
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"duplicate value - row",
			args{
				[]setValueArgs{
					{0, 0, OneRuneValue},
					{0, 8, OneRuneValue},
				},
			},
		},
		{
			"duplicate value - column",
			args{
				[]setValueArgs{
					{0, 0, OneRuneValue},
					{8, 0, OneRuneValue},
				},
			},
		},
		{
			"duplicate value - section",
			args{
				[]setValueArgs{
					{0, 0, OneRuneValue},
					{1, 1, OneRuneValue},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := createEmptyBoard()

			for _, setValueArgs := range tt.args.setValueArgs {
				_, err := board.setValue(setValueArgs.row, setValueArgs.col, setValueArgs.value)
				if err != nil {
					t.Error(err)
				}
			}

			invalid := board.Validate()

			if !invalid.Any() {
				t.Error("expected Invalid.Any() to be true")
			}
		})
	}
}
