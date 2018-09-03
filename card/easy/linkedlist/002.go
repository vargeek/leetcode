package linkedlist

// 删除链表的倒数第N个节点
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/42/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListFromInts create a new list
func NewListFromInts(vals ...int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for _, val := range vals {
		if head == nil {
			head = &ListNode{Val: val, Next: nil}
			tail = head
		} else {
			tail.Next = &ListNode{Val: val, Next: nil}
			tail = tail.Next
		}
	}
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := head
	node := head
	for node != nil {
		if n < 0 {
			prev = prev.Next
		} else {
			n--
		}
		node = node.Next
	}

	if n < 0 {
		if prev.Next != nil {
			prev.Next = prev.Next.Next
		}
	} else if n == 0 {
		head = head.Next
	}

	return head
}
