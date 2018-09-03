package easymath

// Roman to Integer

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/878/
// https://leetcode.com/problems/roman-to-integer/description/

func romanToInt(s string) int {
	sum := 0
	idx := 0
	for idx < len(s) {
		c := s[idx]
		switch c {
		case 'I':
			if idx+1 < len(s) {
				if s[idx+1] == 'V' {
					sum += 4
					idx += 2
					continue
				} else if s[idx+1] == 'X' {
					sum += 9
					idx += 2
					continue
				}
			}
			sum++
		case 'V':
			sum += 5
		case 'X':
			if idx+1 < len(s) {
				if s[idx+1] == 'L' {
					sum += 40
					idx += 2
					continue
				} else if s[idx+1] == 'C' {
					sum += 90
					idx += 2
					continue
				}
			}
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			if idx+1 < len(s) {
				if s[idx+1] == 'D' {
					sum += 400
					idx += 2
					continue
				} else if s[idx+1] == 'M' {
					sum += 900
					idx += 2
					continue
				}
			}
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
		idx++
	}

	return sum
}
