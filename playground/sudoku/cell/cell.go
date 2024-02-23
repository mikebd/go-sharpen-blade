package cell

// Cell is a single cell in a Board.
type Cell struct {
	row, col     int
	value        value
	initialValue value
}

// NewCell is a constructor for the Cell struct.
func NewCell(row, col int, runeValue rune) *Cell {
	value := runeToValue(runeValue)
	return &Cell{row, col, value, value}
}

// Empty returns true if the Cell is empty.
func (c *Cell) Empty() bool {
	return c.value == empty
}

func (c *Cell) Equal(other *Cell) bool {
	return c != nil && other != nil &&
		c.row == other.row &&
		c.col == other.col &&
		c.value == other.value &&
		c.initialValue == other.initialValue
}

// SetValue sets the rune value of the Cell and returns the previous rune value.
func (c *Cell) SetValue(runeValue rune) (previousRuneValue rune) {
	previousRuneValue = c.value.rune()
	c.value = runeToValue(runeValue)
	return
}

// Value returns the rune value of the Cell or NilRuneValue if the Cell is nil.
func (c *Cell) Value() rune {
	if c == nil {
		return NilRuneValue
	}
	return c.value.rune()
}

// Value returns the rune value of the Cell or NilRuneValue if the Cell is nil.
func Value(c *Cell) rune {
	return c.Value()
}
