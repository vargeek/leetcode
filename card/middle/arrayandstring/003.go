package arrayandstring

import (
	"sort"
)

// Group Anagrams
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/778/
// https://leetcode.com/problems/group-anagrams/

// RuneSlice sort.Interface
type RuneSlice []rune

func (s RuneSlice) Len() int           { return len(s) }
func (s RuneSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s RuneSlice) Less(i, j int) bool { return s[i] < s[j] }

func groupKey(str string) string {
	runes := RuneSlice(str)
	sort.Sort(runes)
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		key := groupKey(str)
		groups[key] = append(groups[key], str)
	}

	results := [][]string{}
	for _, group := range groups {
		results = append(results, group)
	}

	return results
}
