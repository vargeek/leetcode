package other

import (
	"sort"
)

// Task Scheduler
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/826/
// https://leetcode.com/problems/task-scheduler/description/

func leastInterval1(tasks []byte, n int) int {
	counts := make([]int, 26)
	for _, task := range tasks {
		counts[task-'A']++
	}
	sort.Ints(counts)

	time := 0
	for counts[25] > 0 {
		i := 0
		for i <= n {
			if counts[25] == 0 {
				break
			}
			if i < 26 && counts[25-i] > 0 {
				counts[25-i]--
			}
			time++
			i++
		}
		sort.Ints(counts)
	}

	return time
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func leastInterval(tasks []byte, n int) int {
	counts := make([]int, 26)
	for _, task := range tasks {
		counts[task-'A']++
	}
	sort.Ints(counts)
	maxVal := counts[25] - 1
	idleSlots := maxVal * n
	for i := 24; i >= 0 && counts[i] > 0; i-- {
		idleSlots -= min(counts[i], maxVal)
	}

	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}
	return len(tasks)
}
