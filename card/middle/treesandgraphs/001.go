package treesandgraphs

// Binary Tree Inorder Traversal
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/786/
// https://leetcode.com/problems/binary-tree-inorder-traversal/

func inorderTraversal(root *TreeNode) []int {
	var loop func(*TreeNode, []int) []int
	loop = func(root *TreeNode, acc []int) []int {
		if root == nil {
			return acc
		}
		acc = loop(root.Left, acc)
		acc = append(acc, root.Val)
		acc = loop(root.Right, acc)
		return acc
	}
	return loop(root, []int{})
}
