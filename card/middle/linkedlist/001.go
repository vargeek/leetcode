package linkedlist

import (
	"fmt"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/783/
// https://leetcode.com/problems/add-two-numbers/description/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNodeFromInts(nums ...int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, num := range nums {
		tail.Next = &ListNode{Val: num, Next: nil}
		tail = tail.Next
	}
	return head.Next
}

func printNodeList(head *ListNode) {
	if head == nil {
		fmt.Println("nil")
		return
	}
	for head != nil {
		fmt.Printf("|%d", head.Val)
		head = head.Next
	}
	fmt.Println("|")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	tail := head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{Val: carry % 10, Next: nil}
		tail = tail.Next
		carry /= 10
	}

	return head.Next
}
