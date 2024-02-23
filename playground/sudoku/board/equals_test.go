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

	for i, row := range b.rows {
		if !row.Equal(other.rows[i]) {
			return false
		}
	}
	for i, col := range b.cols {
		if !col.Equal(other.cols[i]) {
			return false
		}
	}
	for i, sct := range b.scts {
		if !sct.Equal(other.scts[i]) {
			return false
		}
	}

	return true
}
