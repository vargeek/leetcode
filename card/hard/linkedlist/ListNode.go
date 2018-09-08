package linkedlist

import (
	"fmt"
)

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toInts() []int {
	ints := make([]int, 0)
	for l != nil {
		ints = append(ints, l.Val)
		l = l.Next
	}
	return ints
}
func (l *ListNode) String() string {
	text := "👉 "
	for l != nil {
		text += fmt.Sprintf("%d->", l.Val)
		l = l.Next
	}
	return text + "nil"
}

func newListNodeFromInts(vals ...int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, val := range vals {
		tail.Next = &ListNode{Val: val, Next: nil}
		tail = tail.Next
	}
	return head.Next
}
