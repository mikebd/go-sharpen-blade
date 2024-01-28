package cli

import (
	"fmt"
	"go-sharpen-blade/playground/sudoku/game/port"
)

type Console struct {
}

func (c *Console) Render(turns []*port.Turn) {
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
	}
}
