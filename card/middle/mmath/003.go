package mmath

// Excel Sheet Column Number
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/817/

func titleToNumber(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - 'A' + 1)
		sum = sum*26 + digit
	}
	return sum
}
