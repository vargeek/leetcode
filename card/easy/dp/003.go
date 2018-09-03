package dp

// Maximum Subarray
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/566/
// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	max := nums[0]
	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}

		if num < 0 && sum < 0 {
			sum = 0
		}
	}

	return max
}
