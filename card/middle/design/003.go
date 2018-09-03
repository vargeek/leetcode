package design

import (
	"math/rand"
)

// Insert Delete GetRandom O(1)
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/112/design/813/

// RandomizedSet the data structure
type RandomizedSet struct {
	indices  map[int]int
	elements []int
}

// Constructor Initialize your data structure here.
func Constructor() RandomizedSet {
	return RandomizedSet{
		indices:  map[int]int{},
		elements: []int{},
	}
}

// Insert Inserts a value to the set. Returns true if the set did not already contain the specified element.
func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.indices[val]; ok {
		return false
	}
	s.indices[val] = len(s.elements)
	s.elements = append(s.elements, val)
	return true
}

// Remove Removes a value from the set. Returns true if the set contained the specified element.
func (s *RandomizedSet) Remove(val int) bool {
	index, ok := s.indices[val]
	if !ok {
		return false
	}

	lastIndex := len(s.elements) - 1
	lastElement := s.elements[lastIndex]
	s.indices[lastElement] = index
	s.elements[index] = lastElement
	delete(s.indices, val)
	s.elements = s.elements[0:lastIndex]
	return true
}

// GetRandom Get a random element from the set.
func (s *RandomizedSet) GetRandom() int {
	if len(s.elements) == 0 {
		return -1
	}

	// randIdx := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(s.elements))
	// return s.elements[randIdx]
	return s.elements[rand.Intn(len(s.elements))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
