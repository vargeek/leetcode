package sortingandsearching

// Search in Rotated Sorted Array
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/804/
// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

func searchPivot(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < nums[0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
func search(nums []int, target int) int {

	pivot := searchPivot(nums)

	binarySearch := func(left, right int) int {
		lo := left
		hi := right
		for lo <= hi {
			mid := (lo + hi) / 2
			if target <= nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		if left <= lo && lo <= right && nums[lo] == target {
			return lo
		}
		return -1
	}

	idx := binarySearch(0, pivot)
	if idx < 0 && pivot+1 < len(nums) {
		idx = binarySearch(pivot+1, len(nums)-1)
	}
	return idx
}
