package hard

// Sliding Window Maximum
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/837/

// func max(num1, num2 int) int {
// 	if num1 > num2 {
// 		return num1
// 	}
// 	return num2
// }
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	length := len(nums)
	maxLToR := make([]int, length)
	maxRToL := make([]int, length)

	maxLToR[0] = nums[0]
	maxRToL[length-1] = nums[length-1]

	for i := 1; i < length; i++ {
		if i%k == 0 {
			maxLToR[i] = nums[i]
		} else {
			maxLToR[i] = max(maxLToR[i-1], nums[i])
		}
		j := length - i - 1
		if (j+1)%k == 0 {
			maxRToL[j] = nums[j]
		} else {
			maxRToL[j] = max(maxRToL[j+1], nums[j])
		}
	}

	maxs := make([]int, length-k+1)
	for i := 0; i+k <= length; i++ {
		maxs[i] = max(maxRToL[i], maxLToR[i+k-1])
	}

	return maxs
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	maxs := make([]int, len(nums)-k+1)
	deque := make([]int, len(nums))

	left := 0
	right := -1
	for i, num := range nums {
		if left <= right && deque[left] < i-k+1 {
			left++
		}
		for left <= right && num > nums[deque[right]] {
			right--
		}
		right++
		deque[right] = i
		if i-k+1 >= 0 {
			maxs[i-k+1] = nums[deque[left]]
		}
	}

	return maxs
}
