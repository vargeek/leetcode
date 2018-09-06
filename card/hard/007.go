package hard

// Longest Consecutive Sequence
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/833/
// https://leetcode.com/problems/longest-consecutive-sequence/solution/#

func longestConsecutive(nums []int) int {
	exists := make(map[int]bool, len(nums))
	for _, num := range nums {
		exists[num] = true
	}
	longest := 0
	for _, num := range nums {
		length := 0
		val := num
		for exists[val] {
			length++
			exists[val] = false
			val--
		}
		val = num + 1
		for exists[val] {
			length++
			exists[val] = false
			val++
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}
