package easy

// 加一
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/27/
// https://leetcode-cn.com/problems/plus-one/description/

func plusOne(digits []int) []int {
	carry := 1
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] >= 10 {
			carry = digits[i] / 10
			digits[i] %= 10
		} else {
			carry = 0
			break
		}
	}

	if carry > 0 {
		for i := 0; i < length/2; i++ {
			digits[i], digits[length-i-1] = digits[length-i-1], digits[i]
		}
		digits = append(digits, carry)
		length = len(digits)
		for i := 0; i < length/2; i++ {
			digits[i], digits[length-i-1] = digits[length-i-1], digits[i]
		}
	}

	return digits
}
