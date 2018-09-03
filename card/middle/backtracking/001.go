package backtracking

// Letter Combinations of a Phone Number
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/793/
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	charss := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	var loop func(int, []byte, []string) []string
	loop = func(idx int, bytes []byte, acc []string) []string {
		if idx >= len(digits) {
			if len(bytes) > 0 {
				acc = append(acc, string(bytes))
			}
			return acc
		}

		chars := charss[digits[idx]-'2']
		for i := 0; i < len(chars); i++ {
			acc = loop(idx+1, append(bytes, chars[i]), acc)
		}

		return acc
	}

	return loop(0, []byte{}, []string{})
}
