package easystrings

// 验证回文字符串
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/36/
// https://leetcode-cn.com/problems/valid-palindrome/description/

func isPalindrome1(s string) bool {
	left := 0
	right := len(s) - 1
	getChar := func(idx int) byte {
		ch := s[idx]
		if ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'z') {
			return ch
		} else if 'A' <= ch && ch <= 'Z' {
			return ch + 'a' - 'A'
		} else {
			return ' '
		}
	}
	for left < right {
		leftChar := getChar(left)
		for left < right && leftChar == ' ' {
			left++
			leftChar = getChar(left)
		}

		rightChar := getChar(right)
		for left < right && rightChar == ' ' {
			right--
			rightChar = getChar(right)
		}
		if left < right && leftChar != ' ' && rightChar != ' ' && leftChar != rightChar {
			return false
		}
		left++
		right--
	}

	return true
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		chl := s[left]
		if 'A' <= chl && chl <= 'Z' {
			chl += 'a' - 'A'
		}
		if !(('0' <= chl && chl <= '9') || ('a' <= chl && chl <= 'z')) {
			left++
			continue
		}
		chr := s[right]
		if 'A' <= chr && chr <= 'Z' {
			chr += 'a' - 'A'
		}

		if !(('0' <= chr && chr <= '9') || ('a' <= chr && chr <= 'z')) {
			right--
			continue
		}
		if chl != chr {
			return false
		}
		left++
		right--

	}
	return true
}
