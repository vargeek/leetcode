package sortingandsearching

// Sort Colors
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/798/
// https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	idx := 0
	for idx <= right {
		if nums[idx] == 0 {
			nums[left], nums[idx] = nums[idx], nums[left]
			left++
			idx++
		} else if nums[idx] == 2 {
			nums[right], nums[idx] = nums[idx], nums[right]
			right--
		} else {
			idx++
		}
	}
}
