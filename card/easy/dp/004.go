package dp

// House Robber

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/576/
// https://leetcode.com/problems/maximum-subarray/

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func rob(nums []int) int {
	rob := 0
	noRob := 0
	for _, num := range nums {
		nextNoRob := max(rob, noRob)
		nextRob := noRob + num
		rob, noRob = nextRob, nextNoRob
	}
	return max(rob, noRob)
}
