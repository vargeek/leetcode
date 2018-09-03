package dp

import (
	"fmt"
)

// Longest Increasing Subsequence
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/810/
// https://leetcode.com/problems/longest-increasing-subsequence/description/

func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums))
	length := -1
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > length {
			length = dp[i]
		}
	}

	fmt.Println(dp)

	return length + 1
}

func lengthOfLIS(nums []int) int {
	is := make([]int, 0, len(nums))

	for _, num := range nums {
		left, right := 0, len(is)-1
		for left <= right {
			mid := (left + right) / 2
			if num <= is[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left >= len(is) {
			is = append(is, num)
		} else {
			is[left] = num
		}
	}

	return len(is)
}
