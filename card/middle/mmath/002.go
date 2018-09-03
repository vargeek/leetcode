package mmath

// Factorial Trailing Zeroes
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/816/

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}
