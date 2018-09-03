package easymath

// Power of Three
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/745/

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 && n%3 == 0 {
		n /= 3
	}

	return n == 1
}
