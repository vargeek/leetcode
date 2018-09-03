package arrayandstring

// Longest Substring Without Repeating Characters
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/779/

func lengthOfLongestSubstring(s string) int {
	lastIdxs := [128]int{}
	fromIdx := 0
	longest := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if lastIdxs[char] >= fromIdx+1 {
			if i-fromIdx > longest {
				longest = i - fromIdx
			}

			fromIdx = lastIdxs[char]
		}
		lastIdxs[char] = i + 1
	}

	if len(s)-fromIdx > longest {
		longest = len(s) - fromIdx
	}

	return longest
}
