package linkedlist

// Odd Even Linked List
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/784/
// https://leetcode.com/problems/odd-even-linked-list/

func oddEvenList(head *ListNode) *ListNode {
	head1 := &ListNode{}
	head2 := &ListNode{}
	tail1 := head1
	tail2 := head2
	for head != nil {
		tail1.Next = head
		tail1 = tail1.Next
		if head.Next != nil {
			tail2.Next = head.Next
			tail2 = tail2.Next
			head = head.Next.Next
		} else {
			break
		}
	}
	tail1.Next = head2.Next
	tail2.Next = nil

	return head1.Next
}
