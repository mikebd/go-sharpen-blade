package port

import . "go-sharpen-blade/playground/sudoku/cell"

// Renderer is an interface for rendering a game before a turn is made
type Renderer interface {
	Render(turns []*Turn, sections CellPointersArray)
}
