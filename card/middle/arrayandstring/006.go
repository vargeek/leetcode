package arrayandstring

import (
	"math"
)

// Increasing Triplet Subsequence

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/781/
// https://leetcode.com/problems/increasing-triplet-subsequence/description/

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min, mid := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= min {
			min = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}
	}

	return false
}

// 2, 4, -2, -3
