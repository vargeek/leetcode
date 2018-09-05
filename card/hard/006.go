package hard

// First Missing Positive
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/832/
// https://leetcode.com/problems/first-missing-positive/description/

func firstMissingPositive(nums []int) int {
	idx := len(nums) - 1
	for idx >= 0 {
		num := nums[idx]
		for num > 0 && num <= len(nums) && num != idx+1 && nums[idx] != nums[num-1] {
			nums[idx], nums[num-1] = nums[num-1], nums[idx]
			num = nums[idx]
		}
		idx--
	}
	for idx, num := range nums {
		if idx+1 != num {
			return idx + 1
		}
	}

	return len(nums) + 1
}
