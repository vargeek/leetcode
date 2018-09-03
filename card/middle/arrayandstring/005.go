package arrayandstring

// Longest Palindromic Substring
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/780/
// https://leetcode.com/problems/longest-palindromic-substring/
// https://leetcode.com/problems/longest-palindromic-substring/description/

// func substringLength(s string, left int, right int) int {
// 	for 0 <= left && right < len(s) && s[left] == s[right] {
// 		left--
// 		right++
// 	}
// 	return right - left - 1
// }

// func max(num1, num2 int) int {
// 	if num1 >= num2 {
// 		return num1
// 	}
// 	return num2
// }
// func longestPalindrome(s string) string {
// 	start := 0
// 	end := -1
// 	for i := 0; i < len(s); i++ {
// 		len1 := substringLength(s, i, i)
// 		len2 := substringLength(s, i, i+1)
// 		len := max(len1, len2)
// 		if len > end-start+1 {
// 			start = i - (len-1)/2
// 			end = i + len/2
// 		}
// 		// fmt.Println(i, len1, len2, len, start, end)
// 	}
// 	return s[start : end+1]
// }

func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := make([]rune, len(s)*2+1)
	runes[0] = '#'
	for i, r := range s {
		runes[2*i+1] = r
		runes[2*i+2] = '#'
	}
	lens := make([]int, len(runes))
	center := 0
	right := 0
	for i := 0; i < len(runes)-1; i++ {
		iMirror := 2*center - i
		if i < right {
			lens[i] = min(right-i, lens[iMirror])
		}
		for i+lens[i]+1 < len(runes) && i-lens[i]-1 >= 0 && runes[i+lens[i]+1] == runes[i-lens[i]-1] {
			lens[i]++
		}
		if i+lens[i] > right {
			center = i
			right = i + lens[i]
		}
	}
	max, idxOfMax := 0, 0
	for i, l := range lens {
		if l > max {
			max = l
			idxOfMax = i
		}
	}
	idx := (idxOfMax - 1) / 2
	length := max
	return s[idx-(length-1)/2 : idx+length/2+1]
}
