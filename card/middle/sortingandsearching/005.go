package sortingandsearching

// Search for a Range
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/802/
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

func searchRange1(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			break
		}
	}
	if left > right {
		return []int{-1, -1}
	}
	idx := (left + right) / 2
	left = idx
	for left > 0 && nums[left-1] == target {
		left--
	}
	right = idx
	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}
	return []int{left, right}
}

func searchRange(nums []int, target int) []int {
	binarySearch := func(leftMost bool) int {
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] > target || (leftMost && target == nums[mid]) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return left
	}

	leftIdx := binarySearch(true)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return []int{-1, -1}
	}
	rightIdx := binarySearch(false) - 1
	return []int{leftIdx, rightIdx}

}
