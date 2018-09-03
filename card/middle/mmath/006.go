package mmath

// Divide Two Integers
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/820/

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == -1<<31 && divisor == -1) {
		return 1<<31 - 1
	}
	sign := 1
	if (dividend >= 0) != (divisor >= 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	quotient := 0
	for dividend >= divisor {
		bit := 1
		multiple := divisor
		for (multiple << 1) <= dividend {
			multiple <<= 1
			bit <<= 1
		}
		quotient |= bit
		dividend -= multiple
	}
	return quotient * sign
}
