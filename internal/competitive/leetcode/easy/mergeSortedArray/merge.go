// https://leetcode.com/problems/merge-sorted-array

package mergeSortedArray

func merge(nums1 []int, firstEmptyPosition int, nums2 []int, _ int) {
	insertPosition := 0

	for _, insertValue := range nums2 {
		for nums1[insertPosition] <= insertValue && insertPosition < firstEmptyPosition {
			insertPosition++
		}

		// Shift values right in nums1 to make room for insertion
		if insertPosition < firstEmptyPosition {
			for position := firstEmptyPosition; position > insertPosition; position-- {
				nums1[position] = nums1[position-1]
			}
			nums1[insertPosition] = insertValue
		}

		nums1[insertPosition] = insertValue
		firstEmptyPosition++
	}
}
