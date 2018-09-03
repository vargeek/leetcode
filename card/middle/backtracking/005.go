package backtracking

// Subsets
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/796/

func subsets(nums []int) [][]int {
	subset := make([]int, len(nums))
	var loop func(int, int, [][]int) [][]int
	loop = func(step int, idx int, acc [][]int) [][]int {
		if step >= len(nums) {
			tmp := make([]int, idx)
			copy(tmp, subset)
			acc = append(acc, tmp)
			return acc
		}
		subset[idx] = nums[step]
		idx++
		acc = loop(step+1, idx, acc)
		idx--
		acc = loop(step+1, idx, acc)
		return acc
	}

	return loop(0, 0, [][]int{})
}
