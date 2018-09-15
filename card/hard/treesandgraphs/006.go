package treesandgraphs

// Course Schedule
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/847/
// https://leetcode.com/problems/course-schedule/description/

// 20ms
func canFinish(numCourses int, prerequisites [][]int) bool {
	canFinish := true

	suffixes := map[int][]int{}
	for _, pre := range prerequisites {
		suffixes[pre[1]] = append(suffixes[pre[1]], pre[0])
	}

	h := map[int]int{}
	var lookup func(int) int
	lookup = func(c int) int {
		if g, ok := h[c]; ok {
			if g == -1 {
				canFinish = false
				return c
			}
			return g
		}

		h[c] = -1
		suffix := suffixes[c]
		if len(suffix) == 0 {
			h[c] = c
			return c
		}
		for _, suff := range suffix {
			lookup(suff)
		}
		h[c] = h[suffix[0]]
		return h[c]
	}
	for c := range suffixes {
		lookup(c)
	}

	return canFinish
}
