package easystrings

import (
	"fmt"
)

// 数数并说
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/39/
// https://leetcode-cn.com/problems/count-and-say/description/

func countAndSay1(n int) string {
	next := func(str string) string {
		text := ""
		pre := 0
		length := len(str)
		for i := 0; i < length-1; i++ {
			if str[i] != str[i+1] {
				text = fmt.Sprintf("%s%d%c", text, i-pre+1, str[i])
				pre = i + 1
			}
		}
		if length > 0 {
			text = fmt.Sprintf("%s%d%c", text, length-pre, str[length-1])
		}

		return text
	}
	str := "1"
	for i := 1; i < n; i++ {
		str = next(str)
	}

	return str
}

func countAndSay(n int) string {
	bytes := []byte{'1'}

	for i := 1; i < n; i++ {
		bs := []byte{}
		length := len(bytes)
		pre := 0
		for j := 0; j < length-1; j++ {
			if bytes[j] != bytes[j+1] {
				bs = append(bs, '0'+byte(j-pre+1), bytes[j])
				pre = j + 1
			}
		}
		bs = append(bs, '0'+byte(length-pre), bytes[length-1])
		bytes = bs
	}
	return string(bytes)
}
