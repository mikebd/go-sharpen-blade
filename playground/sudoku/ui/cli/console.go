package cli

import (
	"fmt"
	"github.com/inancgumus/screen"
	. "go-sharpen-blade/playground/sudoku/cell"
	"go-sharpen-blade/playground/sudoku/game/port"
)

type Console struct {
}

func (c *Console) Render(turns []*port.Turn, sections CellPointersArray) {
	screen.Clear()
	screen.MoveTopLeft()
	if len(turns) == 0 {
		fmt.Println("Initial Board")

	} else {
		fmt.Println("Board after turn", len(turns))
		lastTurn := *turns[len(turns)-1]
		lastTurnValid := "valid"
		if !lastTurn.Valid() {
			lastTurnValid = "invalid"
		}
		fmt.Printf("\tLast move was %s: row %d, column %d, value %c\n",
			lastTurnValid, lastTurn.Row(), lastTurn.Col(), lastTurn.Value())
		renderBoard(turns, sections)
	}
}
