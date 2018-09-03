package backtracking

// Generate Parentheses
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/794/
// https://leetcode.com/problems/generate-parentheses/description/

func generateParenthesis(n int) []string {
	var loop func([]byte, int, int, []string) []string
	loop = func(bytes []byte, left int, right int, acc []string) []string {
		if left == n && right == n {
			if len(bytes) > 0 {
				acc = append(acc, string(bytes))
			}
			return acc
		}

		if left < n {
			acc = loop(append(bytes, '('), left+1, right, acc)
		}
		if right+1 <= left && right < n {
			acc = loop(append(bytes, ')'), left, right+1, acc)
		}

		return acc
	}

	return loop([]byte{}, 0, 0, []string{})
}
