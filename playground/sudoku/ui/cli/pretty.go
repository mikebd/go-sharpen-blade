package cli

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"go-sharpen-blade/playground/sudoku/game/port"
	"os"
)

func renderBoard(turns []*port.Turn) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Row", "Column", "Value", "Valid"})
	/*	for i, turn := range turns {
			t.AppendRow(table.Row{i, turn.Row(), turn.Col(), turn.Value(), turn.Valid()})
		}
	*/
	t.Render()
}
