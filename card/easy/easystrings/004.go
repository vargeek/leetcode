package easystrings

// 有效的字母异位词
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/35/

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts1 := map[rune]int{}
	counts2 := map[rune]int{}
	for _, c := range s {
		counts1[c]++
	}
	for _, c := range t {
		counts2[c]++
	}
	for i, c1 := range counts1 {
		c2 := counts2[i]
		if c1 != c2 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts1 := [26]int{}
	counts2 := [26]int{}
	for _, c := range s {
		counts1[c-'a']++
	}
	for _, c := range t {
		counts2[c-'a']++
	}
	for i, c1 := range counts1 {
		if c1 != counts2[i] {
			return false
		}
	}
	return true
}
