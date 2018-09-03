package easytree

// Binary Tree Level Order Traversal

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/628/
// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//   3
//  / \
// 9  20
//   /  \
//  15   7
func levelOrder(root *TreeNode) [][]int {
	var travel func(*TreeNode, int, [][]int) [][]int
	travel = func(root *TreeNode, depth int, ans [][]int) [][]int {
		if root == nil {
			return ans
		}
		for len(ans) < depth+1 {
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], root.Val)
		ans = travel(root.Left, depth+1, ans)
		ans = travel(root.Right, depth+1, ans)
		return ans
	}
	return travel(root, 0, [][]int{})
}
