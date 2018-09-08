package linkedlist

// Merge k Sorted Lists
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/839/

func mergeKLists1(lists []*ListNode) *ListNode {
	type ListNodeHead struct {
		list *ListNode
		next *ListNodeHead
	}

	heads := &ListNodeHead{}
	tail := heads
	for _, list := range lists {
		if list != nil {
			tail.next = &ListNodeHead{
				list: list,
				next: nil,
			}
			tail = tail.next
		}
	}

	result := &ListNode{}
	resultTail := result
	for heads.next != nil {
		ahead := heads
		aheadOfMin := ahead
		for ahead.next != nil {
			if ahead.next.list.Val < aheadOfMin.next.list.Val {
				aheadOfMin = ahead
			}
			ahead = ahead.next
		}
		resultTail.Next = aheadOfMin.next.list
		aheadOfMin.next.list = aheadOfMin.next.list.Next
		resultTail.Next.Next = nil
		if aheadOfMin.next.list == nil {
			aheadOfMin.next = aheadOfMin.next.next
		}
		resultTail = resultTail.Next

	}

	return result.Next
}

// 优化
func merge(l, r *ListNode) *ListNode {
	head := ListNode{}
	tail := &head
	for {
		if l == nil {
			tail.Next = r
			break
		}
		if r == nil {
			tail.Next = l
			break
		}
		if l.Val < r.Val {
			tail.Next = l
			l = l.Next
		} else {
			tail.Next = r
			r = r.Next
		}
		tail = tail.Next
	}

	return head.Next
}
func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return merge(lists[0], lists[1])
	default:
		mid := len(lists) / 2
		l := mergeKLists(lists[:mid])
		r := mergeKLists(lists[mid:])
		return merge(l, r)
	}
}
