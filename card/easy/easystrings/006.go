package easystrings

// 字符串转整数（atoi）
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/37/
// https://leetcode-cn.com/problems/string-to-integer-atoi/description/

func myAtoi(str string) int {
	min := -(1 << 31)
	max := (1 << 31) - 1
	ans := 0
	sign := 1
	idx := 0
	length := len(str)
	for idx < length {
		if str[idx] == ' ' {
			idx++
		} else {
			break
		}
	}
	if idx < length {
		if str[idx] == '-' {
			idx++
			sign = -1
		} else if str[idx] == '+' {
			idx++
		}
	}
	for idx < length {
		ch := str[idx]
		if '0' <= ch && ch <= '9' {
			ans = ans*10 + int(ch-'0')*sign
			if ans < min {
				ans = min
				break
			} else if ans > max {
				ans = max
				break
			}
		} else {
			break
		}
		idx++
	}

	return ans
}
