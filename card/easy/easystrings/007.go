package easystrings

// 实现strStr()
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/38/

func strStr(haystack string, needle string) int {
	len1 := len(haystack)
	len2 := len(needle)
	for i := 0; i < len1-len2+1; i++ {
		match := true
		for j := 0; j < len2; j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
