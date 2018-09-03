package linkedlist

// 反转链表
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/43/
// https://leetcode-cn.com/problems/reverse-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	var list *ListNode
	for head != nil {
		list = &ListNode{Val: head.Val, Next: list}
		head = head.Next
	}
	return list
}

func reverseList(head *ListNode) *ListNode {
	// 1 2 3 4 5
	//
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
