package linkedlist

// 合并两个有序链表
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/44/
// https://leetcode-cn.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	tail := head
	var node *ListNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		tail.Next = node
		tail = node
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return head
}
