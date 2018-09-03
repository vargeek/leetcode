package backtracking

// Permutations
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/795/
// https://leetcode.com/problems/permutations/description/

func permute(nums []int) [][]int {
	var loop func(int, [][]int) [][]int
	loop = func(idx int, acc [][]int) [][]int {
		if idx >= len(nums) {
			if len(nums) > 0 {
				tmp := make([]int, len(nums))
				copy(tmp, nums)
				acc = append(acc, tmp)
			}
			return acc
		}

		for i := idx; i < len(nums); i++ {
			nums[i], nums[idx] = nums[idx], nums[i]
			acc = loop(idx+1, acc)
			nums[i], nums[idx] = nums[idx], nums[i]
		}

		return acc
	}
	return loop(0, [][]int{})
}
