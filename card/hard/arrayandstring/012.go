package hard

// Minimum Window Substring
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/838/

func minWindow1(s string, t string) string {
	expects := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		expects[t[i]]++
	}
	totalExpect := len(t)

	counts := map[byte]int{}
	total := 0
	ll, rr := 0, len(s)-1
	l := 0
	for r := 0; r < len(s); r++ {
		chR := s[r]
		if expect, ok := expects[chR]; ok {
			counts[chR]++
			if counts[chR] <= expect {
				total++
			}
		}

		for l <= r {
			chL := s[l]
			if expect, ok := expects[chL]; ok {
				if counts[chL] > expect {
					counts[chL]--
					l++
				} else {
					break
				}
			} else {
				l++
			}
		}
		if total == totalExpect && r-l+1 < rr-ll+1 {
			ll, rr = l, r
		}
	}
	if total == totalExpect {
		return string([]byte(s)[ll : rr+1])
	}

	return ""
}

func minWindow(s string, t string) string {
	counts := [256]int{}
	for i := 0; i < len(t); i++ {
		counts[t[i]]++
	}
	expect := len(t)

	total := 0
	ll, rr := 0, len(s)-1
	l := 0
	for r := 0; r < len(s); r++ {
		chR := s[r]
		counts[chR]--
		if counts[chR] >= 0 {
			total++
		}

		for l <= r {
			chL := s[l]
			if counts[chL] < 0 {
				counts[chL]++
				l++
			} else {
				break
			}
		}
		if total == expect && r-l+1 < rr-ll+1 {
			ll, rr = l, r
		}
	}
	if total == expect {
		return string([]byte(s)[ll : rr+1])
	}

	return ""
}
