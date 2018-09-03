package easytree

// Symmetric Tree
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/627/
// https://leetcode.com/problems/symmetric-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var isMirror func(*TreeNode, *TreeNode) bool
	isMirror = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && isMirror(left.Right, right.Left) && isMirror(left.Left, right.Right)
	}
	return root == nil || isMirror(root.Left, root.Right)
}
