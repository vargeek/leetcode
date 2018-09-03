package main

import (
	"fmt"
)

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var preNode *ListNode
	high := 0
	for l1 != nil || l2 != nil {
		current := ListNode{Val: high, Next: nil}
		if l1 != nil {
			current.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current.Val += l2.Val
			l2 = l2.Next
		}
		high = current.Val / 10
		current.Val %= 10
		if head == nil {
			head = &current
		}
		if preNode != nil {
			preNode.Next = &current
		}
		preNode = &current
	}
	if high != 0 {
		current := ListNode{Val: high, Next: nil}
		if preNode != nil {
			preNode.Next = &current
		}
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	high := 0
	for l1 != nil || l2 != nil || high != 0 {
		sum := high
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			current.Next = &ListNode{Val: sum % 10, Next: nil}
			high = sum / 10
		} else {
			current.Next = &ListNode{Val: sum, Next: nil}
			high = 0
		}
		current = current.Next
	}
	return head.Next
}
func printListNode(list *ListNode) {
	output := ""
	for list != nil {
		output = output + fmt.Sprintf("|%d", list.Val)
		list = list.Next
	}
	if len(output) == 0 {
		fmt.Println("nil")
	} else {
		fmt.Printf("%s|\n", output)
	}
}

func main() {
	node1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	node2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	result := addTwoNumbers(&node1, &node2)
	printListNode(&node1)
	printListNode(&node2)
	printListNode(result)

}
