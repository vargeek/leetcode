package easy

import (
	"fmt"
)

// 从排序数组中删除重复项
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/21/
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates1(nums []int) int {
	idx := 0
	for _, num := range nums {
		if idx == 0 || num != nums[idx-1] {
			nums[idx] = num
			idx++
		}
	}
	return idx
}
func removeDuplicates(nums []int) int {
	count := len(nums)
	dupCount := 0
	for i := 1; i < count; i++ {
		if nums[i] != nums[i-1] {
			nums[i-dupCount] = nums[i]
		} else {
			dupCount++
		}
	}

	return count - dupCount
}

func main() {
	fmt.Println("xxxx")
}
