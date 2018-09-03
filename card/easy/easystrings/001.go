package easystrings

// 反转字符串
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/32/
// https://leetcode-cn.com/problems/reverse-string/description/

func reverseString1(s string) string {
	bytes := []byte(s)
	length := len(bytes)
	for i := 0; i < length/2; i++ {
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
	return string(bytes)
}

func reverseString(s string) string {
	length := len(s)
	runes := make([]rune, length)
	for i, c := range s {
		runes[length-i-1] = c
	}
	return string(runes)
}
