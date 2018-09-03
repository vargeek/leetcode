package treesandgraphs

// Kth Smallest Element in a BST
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/790/
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

func kthSmallest(root *TreeNode, k int) int {
	var loop func(*TreeNode, int) (int, int)
	loop = func(root *TreeNode, rank int) (int, int) {
		if root == nil {
			return rank, 0
		}
		leftk, leftVal := loop(root.Left, rank)
		if leftk == k {
			return k, leftVal
		}
		if leftk+1 == k {
			return k, root.Val
		}

		right, rightVal := loop(root.Right, leftk+1)
		return right, rightVal
	}
	_, val := loop(root, 0)

	return val
}
