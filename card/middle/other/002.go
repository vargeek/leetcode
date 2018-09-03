package other

import "strconv"

// Evaluate Reverse Polish Notation

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/823/
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		l := len(stack)
		switch token {
		case "+":
			stack[l-2] = stack[l-2] + stack[l-1]
			stack = stack[:l-1]
		case "-":
			stack[l-2] = stack[l-2] - stack[l-1]
			stack = stack[:l-1]
		case "*":
			stack[l-2] = stack[l-2] * stack[l-1]
			stack = stack[:l-1]
		case "/":
			stack[l-2] = stack[l-2] / stack[l-1]
			stack = stack[:l-1]
		default:
			if num, err := strconv.Atoi(token); err == nil {
				stack = append(stack, num)
			}
		}
	}

	return stack[0]
}
