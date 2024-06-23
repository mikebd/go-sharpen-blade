package canWin

// Given:
//   - an array where each value is how far you can jump from that index
//   - a starting index within that array
//
// Return:
//
//	True if you can reach an index with a jump size of 0 by starting at that
//	index and repeatedly jumping left or right, false otherwise.
//
// For example, given the array [1 0 2 4], a starting index of:
//
//	0 returns true: you can reach 0 by jumping right one spot.
//	1 returns true: that value is 0.
//	2 returns true: you can jump left two spots and then right one spot.
//	3 returns false: there is no route to reach a 0 jumping left or right 4 spots.
func canWin(board []int, startingIndex int) bool {
	// visited is a slice of bool that tracks which indices we've already
	// visited. This is necessary to prevent infinite recursion.
	visited := make([]bool, len(board))

	return jump(board, startingIndex, visited)
}

func jump(board []int, startingIndex int, visited []bool) bool {
	if outOfBounds(board, startingIndex) {
		return false
	}

	visited[startingIndex] = true

	currentValue := board[startingIndex]
	if currentValue == 0 {
		return true
	}

	leftIndex := startingIndex - currentValue
	rightIndex := startingIndex + currentValue
	return isWinningIndex(board, leftIndex, visited) ||
		isWinningIndex(board, rightIndex, visited)
}

func isWinningIndex(board []int, index int, visited []bool) bool {
	indexAvailable := !outOfBounds(board, index)
	if indexAvailable && !visited[index] {
		return jump(board, index, visited)
	}
	return false
}

func outOfBounds(board []int, index int) bool {
	return index < 0 || index >= len(board)
}
