package backtracking

// Palindrome Partitioning

// https://leetcode.com/explore/interview/card/top-interview-questions-hard/119/backtracking/852/

func partition(s string) [][]string {

	result := make([][]string, 0)

	var loop func(int, []string)
	loop = func(left int, subs []string) {
		if left >= len(s) {
			p := make([]string, len(subs))
			copy(p, subs)
			result = append(result, p)
			return
		}

		for right := len(s) - 1; right >= left; right-- {
			mid := (left + right + 1) / 2
			isPalindrome := true
			for i := left; i < mid; i++ {
				if s[i] != s[right-(i-left)] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome {
				subs = append(subs, string(s[left:right+1]))
				loop(right+1, subs)
				subs = subs[:len(subs)-1]
			}
		}
	}
	loop(0, make([]string, 0, len(s)))

	return result
}
