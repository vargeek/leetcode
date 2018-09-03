package easytree

// Validate Binary Search Tree

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/625/
// https://leetcode.com/problems/validate-binary-search-tree/description/

func isValidBST(root *TreeNode) bool {
	var isBST func(*TreeNode, int, int) bool
	isBST = func(root *TreeNode, min int, max int) bool {
		return root == nil || (root.Val > min && root.Val < max && isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max))
	}
	return isBST(root, -(1 << 32), 1<<32)
}
