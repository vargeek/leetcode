package design

import (
	"testing"
)

func TestMinStackConstructor(t *testing.T) {
	// MinStack minStack = new MinStack();
	// minStack.push(-2);
	// minStack.push(0);
	// minStack.push(-3);
	// minStack.getMin();   --> Returns -3.
	// minStack.pop();
	// minStack.top();      --> Returns 0.
	// minStack.getMin();   --> Returns -2.

	// ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

	// [null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483647,null,-2147483648,-2147483648,null,2147483647]

	s := MinStackConstructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 {
		t.Errorf("min value should be %d\n", -3)
	}
	s.Pop()
	if s.Top() != 0 {
		t.Errorf("top value should be %d\n", 0)
	}
	if s.GetMin() != -2 {
		t.Errorf("top value should be %d\n", -2)
	}

}
