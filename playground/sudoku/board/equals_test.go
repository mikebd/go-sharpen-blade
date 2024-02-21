package board

// equals returns true if the two boards have equal current values.
// equals is only meant to be used for testing, it does not
// have a parallelized implementation.
func (board *Board) equals(other *Board) bool {
	if board == nil || other == nil {
		return false
	}

	if board.Size() != other.Size() {
		return false
	}

	for row := 0; row < board.Size(); row++ {
		for col := 0; col < board.Size(); col++ {
			if board.rows[row][col].Value() != other.rows[row][col].Value() {
				return false
			}
		}
	}

	return true
}
