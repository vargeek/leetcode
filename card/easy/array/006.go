package easy

import (
	"fmt"
	"sort"
)

// 两个数组的交集 II
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/26/
// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/

func intersect(nums1 []int, nums2 []int) []int {
	ans := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return ans
}

type bstNode struct {
	val   int
	count int
	left  *bstNode
	right *bstNode
}

func (node *bstNode) insert(otherNode *bstNode) {
	if otherNode == nil {
		return
	}

	if node.val == otherNode.val {
		node.val += otherNode.val
		if node.left == nil {
			node.left = otherNode.left
		} else {
			node.left.insert(otherNode.left)
		}
		if node.right == nil {
			node.right = otherNode.right
		} else {
			node.right.insert(otherNode.right)
		}
	} else if otherNode.val < node.val {
		if node.left == nil {
			node.left = otherNode
		} else {
			node.left.insert(otherNode)
		}
	} else {
		if node.right == nil {
			node.right = otherNode
		} else {
			node.right.insert(otherNode)
		}
	}
}

func (node *bstNode) print() {
	fmt.Printf("|%d:%d", node.val, node.count)
	if node.left != nil {
		node.left.print()
	}
	if node.right != nil {
		node.right.print()
	}
}

type bstTree struct {
	node *bstNode
}

func newBSTTreeFromInts(vals []int) *bstTree {
	tree := &bstTree{}
	for _, val := range vals {
		tree.add(val)
	}
	return tree
}
func (tree *bstTree) add(val int) {
	node := &tree.node
	for (*node) != nil {
		if (*node).val == val {
			(*node).count++
			return
		} else if val < (*node).val {
			node = &(*node).left
		} else {
			node = &(*node).right
		}
	}
	*node = &bstNode{val: val, count: 1, left: nil, right: nil}
}

func (tree *bstTree) contains(val int) bool {
	node := tree.node
	for node != nil {
		if node.val == val {
			return true
		} else if val < node.val {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

func (tree *bstTree) delete(val int) {
	node := &tree.node
	for (*node) != nil {
		if (*node).val == val {
			(*node).count--
			if (*node).count <= 0 {
				right := (*node).right
				left := (*node).left
				if right != nil {
					(*node) = right
					right.insert(left)
				} else {
					(*node) = left
				}
			}
			return
		} else if val < (*node).val {
			node = &(*node).left
		} else {
			node = &(*node).right
		}
	}
}

func (tree *bstTree) print() {
	if tree.node != nil {
		tree.node.print()
		fmt.Println("|")
	} else {
		fmt.Println("empty tree")
	}
}

func intersect1(nums1 []int, nums2 []int) []int {
	ans := []int{}

	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	sort.Ints(nums1)
	tree := newBSTTreeFromInts(nums1)

	for _, val := range nums2 {
		if tree.contains(val) {
			ans = append(ans, val)
			tree.delete(val)
		}
	}

	return ans
}
