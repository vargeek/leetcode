package easy

// 只出现一次的数字
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/25/
// https://leetcode-cn.com/problems/single-number/description/

func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
