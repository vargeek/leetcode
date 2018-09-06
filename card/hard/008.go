package hard

// Find the Duplicate Number
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/834/
// https://leetcode.com/problems/find-the-duplicate-number/description/

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	// Find the "entrance" to the cycle.
	ptr1 := nums[0]
	ptr2 := slow
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}
