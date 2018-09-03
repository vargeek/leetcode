package easytree

// Maximum Depth of Binary Tree
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
