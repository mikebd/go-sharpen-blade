package cli

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	. "go-sharpen-blade/playground/sudoku/cell"
	"go-sharpen-blade/playground/sudoku/game/port"
	"os"
)

func renderBoard(turns []*port.Turn, rows CellPointersArray) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Style().Color.IndexColumn = text.Colors{text.BgBlue, text.FgHiWhite}
	t.AppendHeader(table.Row{"#", "A", "B", "C", "D", "E", "F", "G", "H", "I"})
	t.AppendRow(table.Row{"1", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"2", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"3", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"4", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"5", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"6", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"7", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"8", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	t.AppendRow(table.Row{"9", "1", "2", "3", "4", "5", "6", "7", "8", "9"})

	/*	for i, turn := range turns {
			t.AppendRow(table.Row{i, turn.Row(), turn.Col(), turn.Value(), turn.Valid()})
		}
	*/
	t.Render()
}
