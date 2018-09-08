package hard

// Product of Array Except Self
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/827/

func productExceptSelf1(nums []int) []int {
	var loop func(int, int, []int) int
	loop = func(step, acc int, output []int) int {
		if step == len(nums) {
			return 1
		}
		suffix := loop(step+1, acc*nums[step], output)
		output[step] = acc * suffix
		return nums[step] * suffix
	}
	output := make([]int, len(nums))
	loop(0, 1, output)
	return output
}

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	p := 1
	for i, num := range nums {
		output[i] = p
		p *= num
	}

	p = 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= p
		p *= nums[i]
	}
	return output
}
