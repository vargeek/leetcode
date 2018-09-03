package mmath

// Happy Number
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/815/

func isHappy(n int) bool {
	getSum := func(num int) int {
		sum := 0
		for num != 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	slow, fast := n, n
	for {
		slow = getSum(slow)
		fast = getSum(getSum(fast))
		if slow == 1 || fast == 1 || slow == fast {
			break
		}
	}

	return slow == 1 || fast == 1
}
