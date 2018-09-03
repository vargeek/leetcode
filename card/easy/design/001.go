package design

import (
	"math/rand"
)

// Shuffle an Array
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/670/
// https://leetcode.com/problems/shuffle-an-array/description/

// Solution contains an array
type Solution struct {
	nums   []int
	backup []int
}

// Constructor return a new Socution struct
func Constructor(nums []int) Solution {
	backup := make([]int, len(nums))
	copy(backup, nums)

	return Solution{nums, backup}
}

// Reset Resets the array to its original configuration and return it.
func (s *Solution) Reset() []int {
	copy(s.nums, s.backup)
	return s.nums
}

// Shuffle Returns a random shuffling of the array.
func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n-1; i++ {
		randIdx := rand.Intn(n-i) + i
		s.nums[randIdx], s.nums[i] = s.nums[i], s.nums[randIdx]
	}
	return s.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
