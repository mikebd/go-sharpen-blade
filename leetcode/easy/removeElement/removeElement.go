// https://leetcode.com/problems/remove-element

package removeElement

func removeElement(nums []int, val int) int {
	nonMatches := 0
	writePosition := 0

	for _, number := range nums {
		if number != val {
			nonMatches++
			nums[writePosition] = number
			writePosition++
		}
	}

	return nonMatches
}
