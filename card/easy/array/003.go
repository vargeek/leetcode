package easy

// 旋转数组
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/23/
// https://leetcode-cn.com/problems/rotate-array/description/

func rotate1(nums []int, k int) {
	count := len(nums)
	if count == 0 {
		return
	}

	for t := 0; t < k; t++ {
		last := nums[count-1]
		for i := count - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = last
	}
}

func rotate(nums []int, k int) {
	count := len(nums)
	if count == 0 {
		return
	}

	k %= count

	idx := 0
	bak := nums[idx]
	bakIdx := idx
	for i := 0; i < count; i++ {
		leftIdx := idx - k
		if leftIdx < 0 {
			leftIdx += count
		}
		if leftIdx == bakIdx {
			nums[idx] = bak
			if leftIdx > 0 {
				idx = leftIdx - 1
			} else {
				idx = count - 1
			}
			bakIdx = idx
			bak = nums[bakIdx]
		} else {
			nums[idx] = nums[leftIdx]
			idx = leftIdx
		}
	}
}

func rotate2(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}
	k %= length
	k = length - k
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}
	for i := 0; i < (length-k)/2; i++ {
		nums[k+i], nums[length-i-1] = nums[length-i-1], nums[k+i]
	}
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}
