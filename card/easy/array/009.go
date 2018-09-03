package easy

import "sort"

// 两数之和
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/29/
// https://leetcode-cn.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	count := len(nums)
	tmpNums := make([]int, count)
	copy(tmpNums, nums)
	sort.Ints(tmpNums)

	left := 0
	right := count - 1
	for left < right {
		sum := tmpNums[left] + tmpNums[right]
		if sum > target {
			right = right - 1
		} else if sum < target {
			left = left + 1
		} else {
			break
		}
	}

	res := make([]int, 0, 2)
	if left < right {
		for i, v := range nums {
			if v == tmpNums[left] || v == tmpNums[right] {
				res = append(res, i)
			}
		}
	}
	return res
}
