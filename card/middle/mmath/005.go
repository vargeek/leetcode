package mmath

// Sqrt(x)
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/819/

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		if x <= mid*mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left*left > x {
		return right
	}
	return left
}
