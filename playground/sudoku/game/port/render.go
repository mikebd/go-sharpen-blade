package port

// Render is an interface for rendering a game before a turn is made
type Renderer interface {
	Render(turns []*Turn)
}
