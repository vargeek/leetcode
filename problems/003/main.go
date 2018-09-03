package main

func lengthOfLongestSubstring1(s string) int {
	flags := make(map[byte]bool)
	bytes := []byte(s)
	first := 0
	last := 0
	count := len(bytes)
	maxLength := 0
	for last < count {
		lastVal := bytes[last]
		last++
		if !flags[lastVal] {
			flags[lastVal] = true
			if last-first > maxLength {
				maxLength = last - first
			}
		} else {
			for first < last-1 {
				firstVal := bytes[first]
				first++
				if firstVal == lastVal {
					break
				} else {
					flags[firstVal] = false
				}
			}
		}
	}
	return maxLength
}

func lengthOfLongestSubstring2(s string) int {
	maxLength := 0
	currentLength := 0
	lastIndex := [128]int{}
	for idx, val := range s {
		// 当前长度+1
		// 字母上一次出现位置到当前位置的长度
		currentLength++
		tmp := idx - lastIndex[val] + 1
		if tmp < currentLength {
			currentLength = tmp
		}
		lastIndex[val] = idx + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	currentLength := 0
	lastIndex := [128]int{}
	last := 1
	for idx, val := range s {
		// 未被重复字符切断，长度+1
		if lastIndex[val] < last {
			currentLength++
		} else {
			// 已经被切断，更新开头位置和最大长度
			last = lastIndex[val]
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = idx - last + 1
		}

		lastIndex[val] = idx + 1
	}
	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}

func assert(condition bool) {
	if !condition {
		panic("测试不通过")
	}
}

func main() {
	assert(lengthOfLongestSubstring("abcabcabc") == 3)
	assert(lengthOfLongestSubstring("cbaabcabc") == 3)
	assert(lengthOfLongestSubstring("bbbbb") == 1)

	assert(lengthOfLongestSubstring("pwwkew") == 3)

	assert(lengthOfLongestSubstring("cdd") == 2)

	assert(lengthOfLongestSubstring(" ") == 1)

}
