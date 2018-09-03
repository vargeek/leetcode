package easy

// 移动零
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/28/
// https://leetcode-cn.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	length := len(nums)
	idx := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	for i := idx; i < length; i++ {
		nums[i] = 0
	}
}
