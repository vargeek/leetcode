package easystrings

// 最长公共前缀
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/40/
// https://leetcode-cn.com/problems/longest-common-prefix/description/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	idx := 0
outerLoop:
	for {
		ch := byte(0)
		for i := 0; i < len(strs); i++ {
			if idx >= len(strs[i]) {
				break outerLoop
			}
			if i == 0 {
				ch = strs[i][idx]
			} else {
				if strs[i][idx] != ch {
					break outerLoop
				}
			}
		}
		idx++
	}

	return string(strs[0][:idx])
}
