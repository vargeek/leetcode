package sortingandsearching

import (
	"sort"
)

// Merge Intervals
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/803/
// https://leetcode.com/problems/merge-intervals/

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

// IntervalSlice [] Interval
type IntervalSlice []Interval

func (s IntervalSlice) Len() int {
	return len(s)
}
func (s IntervalSlice) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}
func (s IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals []Interval) []Interval {
	sort.Sort(IntervalSlice(intervals))
	ans := []Interval{}
	var curr *Interval
	for i := 0; i < len(intervals); i++ {
		if curr == nil {
			curr = &Interval{Start: intervals[i].Start, End: intervals[i].End}
		} else if intervals[i].End > curr.End {
			curr.End = intervals[i].End
		}
		if i+1 >= len(intervals) || intervals[i+1].Start > curr.End {
			ans = append(ans, *curr)
			curr = nil
		}
	}

	return ans
}
