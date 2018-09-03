package others

// Valid Parentheses
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/721/
// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		lastChar := ' '
		if len(stack) > 0 {
			lastChar = stack[len(stack)-1]
		}

		if (char == ')' && lastChar == '(') || (char == '}' && lastChar == '{') || (char == ']' && lastChar == '[') {
			stack = stack[:len(stack)-1]
		} else if char == ')' || char == '}' || char == ']' {
			return false
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
