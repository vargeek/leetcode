package linkedlist

// Sort List
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/840/

func sortList(head *ListNode) *ListNode {
	var loop func(*ListNode) (*ListNode, *ListNode)
	loop = func(head *ListNode) (*ListNode, *ListNode) {
		if head == nil || head.Next == nil {
			return head, head
		}
		headL, headM, headR := ListNode{}, ListNode{}, ListNode{}
		tailL, tailM, tailR := &headL, &headM, &headR
		// val可以考虑使用最大值和最小值的平均值，避免输入链表已经排好序时，效率下降
		val := head.Val
		for head != nil {
			if head.Val == val {
				tailM.Next = head
				tailM = tailM.Next
			} else if head.Val < val {
				tailL.Next = head
				tailL = tailL.Next
			} else {
				tailR.Next = head
				tailR = tailR.Next
			}
			head = head.Next
		}
		tailL.Next = nil
		tailM.Next = nil
		tailR.Next = nil

		hl, tl := loop(headL.Next)
		hr, tr := loop(headR.Next)
		if hl != nil {
			tl.Next = headM.Next
		} else {
			hl = headM.Next
		}

		tailM.Next = hr
		if hr == nil {
			tr = tailM
		}
		return hl, tr
	}
	h, _ := loop(head)
	return h
}
