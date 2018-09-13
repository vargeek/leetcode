package treesandgraphs

// Binary Tree Maximum Path Sum
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/845/

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxVal := root.Val

	var loop func(*TreeNode) int
	loop = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := loop(root.Left)
		right := loop(root.Right)
		if left < 0 {
			left = 0
		}
		if right < 0 {
			right = 0
		}
		if left+root.Val+right > maxVal {
			maxVal = left + root.Val + right
		}
		return root.Val + max3(left, right, 0)

	}
	loop(root)

	return maxVal
}
