package easystrings

// 字符串中的第一个唯一字符
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/34/
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar1(s string) int {
	counts := [128]int{}
	for _, c := range s {
		counts[c]++
	}
	for idx, c := range s {
		if counts[c] == 1 {
			return idx
		}
	}
	return -1
}

func firstUniqChar(s string) int {
	counts := [26]int{}
	for _, c := range s {
		counts[c-'a']++
	}
	for idx, c := range s {
		if counts[c-'a'] == 1 {
			return idx
		}
	}
	return -1
}
