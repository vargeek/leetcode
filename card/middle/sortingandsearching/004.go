package sortingandsearching

// Find Peak Element
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/801/

func findPeakElement(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if (i == 0 || nums[i-1] < nums[i]) && (i == len(nums)-1 || nums[i+1] < nums[i]) {
			return i
		}
	}

	return 0
}
