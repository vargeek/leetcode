package dp

// Climbing Stairs
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/569/

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	step0 := 1
	step1 := 1
	step := 0
	for i := 1; i < n; i++ {
		step = step0 + step1
		step0, step1 = step1, step
	}
	return step
}
