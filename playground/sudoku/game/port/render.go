package port

import . "go-sharpen-blade/playground/sudoku/cell"

// Renderer is an interface for rendering a game before a turn is made
type Renderer interface {
	// Overwrites returns true if each Render call overwrites the previous Render call.
	// If so, Render calls may be omitted until the game ends.
	Overwrites() bool

	Render(turns []*Turn, sections CellPointersArray)
}
