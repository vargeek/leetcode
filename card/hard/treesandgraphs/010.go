package treesandgraphs

import (
	"sort"
)

// Count of Smaller Numbers After Self
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/851/

func countSmaller1(nums []int) []int {
	counts := make([]int, len(nums))

	binarySearch := func(left, right, val int) int {
		for left <= right {
			mid := (left + right) / 2
			if val >= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return right
	}

	right := len(nums) - 1
	for i := right; i >= 0; i-- {
		val := nums[i]
		idx := binarySearch(i+1, right, val-1)
		for j := i; j < idx; j++ {
			nums[j] = nums[j+1]
		}
		nums[idx] = val
		counts[i] = right - idx
	}

	return counts
}

// Fenwick tree
type Fenwick struct {
	BIT []int
}

// Sum sum
func (f *Fenwick) Sum(idx int) int {
	sum := 0
	for idx > 0 {
		sum += f.BIT[idx]
		// 父节点
		idx -= idx & (-idx)
	}
	return sum
}

// Inc inc
func (f *Fenwick) Inc(idx int) {
	for idx < len(f.BIT) {
		f.BIT[idx]++
		// 扩大区间
		idx += idx & (-idx)
	}
}

func countSmaller(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	sorted := make([]int, length)
	copy(sorted, nums)
	sort.Ints(sorted)

	ranks := make(map[int]int, length)
	prev := sorted[0] - 1
	for i, val := range sorted {
		if val != prev {
			ranks[val] = i + 1
		}
		prev = val
	}

	result := make([]int, length)
	f := &Fenwick{make([]int, length)}

	for i := length - 1; i >= 0; i-- {
		val := nums[i]
		rank := ranks[val]
		sum := f.Sum(rank - 1)
		f.Inc(rank)
		result[i] = sum
	}

	return result

}
