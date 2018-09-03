package sortingandsearching

// Kth Largest Element in an Array
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/800/
// https://leetcode.com/problems/kth-largest-element-in-an-array/
// [Quickselect](https://en.wikipedia.org/wiki/Quickselect)

func findKthLargest(nums []int, k int) int {
	var loop func(int, int) int
	loop = func(left, right int) int {
		if right <= left {
			return nums[k-1]
		}
		i, j, val := left, right, nums[(left+right)/2]
		for i <= j {
			for i <= j && nums[i] > val {
				i++
			}
			for i <= j && val > nums[j] {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if k-1 <= j {
			return loop(left, j)
		}
		if k-1 >= i {
			return loop(i, right)
		}
		return nums[k-1]
	}

	return loop(0, len(nums)-1)
}
