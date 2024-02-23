package cli

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	. "go-sharpen-blade/playground/sudoku/cell"
	"go-sharpen-blade/playground/sudoku/game/port"
	"os"
)

func renderBoard(turns []*port.Turn, sections CellPointersArray) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Style().Color.IndexColumn = text.Colors{text.BgBlue, text.FgHiWhite}
	//t.AppendHeader(table.Row{"#", "A", "B", "C", "D", "E", "F", "G", "H", "I"})

	for _, row := range sections {
		tableRow := table.Row{}
		for _, cell := range row {
			tableRow = append(tableRow, string(cell.Value()))
		}
		t.AppendRow(tableRow)
	}
	/*	for i, turn := range turns {
			t.AppendRow(table.Row{i, turn.Row(), turn.Col(), turn.Value(), turn.Valid()})
		}
	*/
	t.Render()
}
