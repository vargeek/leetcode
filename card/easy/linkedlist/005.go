package linkedlist

// 回文链表

// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/45/
// https://leetcode-cn.com/problems/palindrome-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var reverseHead *ListNode
	node := head
	for node != nil {
		reverseHead = &ListNode{Val: node.Val, Next: reverseHead}
		node = node.Next
	}

	for head != nil {
		if head.Val != reverseHead.Val {
			return false
		}
		head = head.Next
		reverseHead = reverseHead.Next
	}

	return true
}
