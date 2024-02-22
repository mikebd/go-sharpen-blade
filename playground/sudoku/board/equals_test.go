package board

// equals returns true if the two boards have equal current values.
// equals is only meant to be used for testing, it does not
// have a parallelized implementation.
func (b *board) equals(other *board) bool {
	if b == nil || other == nil {
		return false
	}

	if b.Size() != other.Size() {
		return false
	}

	for row := 0; row < b.Size(); row++ {
		for col := 0; col < b.Size(); col++ {
			if b.rows[row][col].Value() != other.rows[row][col].Value() {
				return false
			}
		}
	}

	return true
}
