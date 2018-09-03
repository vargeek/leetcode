package sortingandsearching

import (
	"sort"
)

// Top K Frequent Elements
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/799/

type numAndFreq struct {
	num  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	pairs := make([]numAndFreq, len(freq))
	idx := 0
	for k, v := range freq {
		pairs[idx].num = k
		pairs[idx].freq = v
		idx++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = pairs[i].num
	}

	return ans
}
