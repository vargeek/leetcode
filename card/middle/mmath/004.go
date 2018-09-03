package mmath

// Pow(x, n)
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/818/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if x == 0.0 {
		return 0.0
	}
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	result := 1.0
	for n > 1 {
		if n&0x01 == 0x01 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result * x
}
