package other

// Majority Element
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/824/

func majorityElement1(nums []int) int {
	counts := map[int]int{}
	nDiv2 := len(nums) / 2
	for _, num := range nums {
		counts[num]++
		if counts[num] > nDiv2 {
			return num
		}
	}
	return 0
}

func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
			count++
		} else if majority == nums[i] {
			count++
		} else {
			count--
		}
	}
	return majority
}
