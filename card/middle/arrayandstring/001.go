package arrayandstring

import (
	"sort"
)

// 3Sum
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/

func threeSum1(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		remain := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right] == remain {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if nums[left]+nums[right] < remain {
				left++
			} else {
				right--
			}
		}

	}

	return result
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		remain := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right] == remain {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if nums[left]+nums[right] < remain {
				left++
			} else {
				right--
			}
		}

	}

	return result
}
