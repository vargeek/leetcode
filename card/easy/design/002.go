package design

// Min Stack
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/562/
// https://leetcode.com/problems/min-stack/

// MinStack min stack
type MinStack struct {
	stack  []int
	minIdx int
}

// MinStackConstructor initialize your data structure here.
// Constructor
func MinStackConstructor() MinStack {
	return MinStack{[]int{}, -1}
}

// Push an element
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if s.minIdx >= 0 && x < s.stack[s.minIdx] {
		s.minIdx = len(s.stack) - 1
	}
}

// Pop an element
func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	if s.minIdx >= len(s.stack) {
		s.minIdx = -1
	}
}

// Top element
func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

// GetMin retrieving the minimum element
func (s *MinStack) GetMin() int {
	if len(s.stack) == 0 {
		return 0
	}
	if s.minIdx < 0 {
		min := s.stack[0]
		s.minIdx = 0
		for i := 1; i < len(s.stack); i++ {
			if s.stack[i] < min {
				min, s.minIdx = s.stack[i], i
			}
		}
	}
	return s.stack[s.minIdx]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
