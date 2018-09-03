package others

import (
	"sort"
)

// Missing Number
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/722/

func missingNumber1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}

func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return (1+len(nums))*len(nums)/2 - sum
}
