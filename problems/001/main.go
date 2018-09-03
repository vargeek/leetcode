package main

import (
	"fmt"
	"sort"
)

func twoSums1(nums []int, target int) []int {
	count := len(nums)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSums(nums []int, target int) []int {
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

func main() {
	target := 4
	nums := []int{10, 9, 6, 5, 7, 8, 4, 3, 2, 1}

	result := twoSums(nums, target)
	fmt.Printf("len: %d\n", len(result))
	if len(result) == 2 {
		fmt.Printf("nums[%d]+nums[%d]=%d\n", result[0], result[1], target)
	}
}
