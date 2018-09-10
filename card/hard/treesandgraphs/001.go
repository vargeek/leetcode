package treesandgraphs

// Word Ladder
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/842/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	reachable := func(a, b string) bool {
		count := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
			if count > 1 {
				return false
			}
		}
		return count == 1
	}

	step := 1
	idx := 0
	count := len(wordList)
	reached := map[string]bool{beginWord: true}
	for idx < count {
		nexts := make(map[string]bool)
		for word := range reached {
			for i := idx; i < count; i++ {
				if reachable(word, wordList[i]) {
					if wordList[i] == endWord {
						return step + 1
					}
					nexts[wordList[i]] = true
					wordList[idx], wordList[i] = wordList[i], wordList[idx]
					idx++
				}
			}
		}

		if len(nexts) == 0 {
			return 0
		}
		reached = nexts
		step++
	}

	return 0
}
