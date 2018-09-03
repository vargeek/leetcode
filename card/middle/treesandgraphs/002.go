package treesandgraphs

// Binary Tree Zigzag Level Order Traversal
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/787/
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var loop func(*TreeNode, int, [][]int) [][]int
	loop = func(root *TreeNode, level int, acc [][]int) [][]int {
		if root == nil {
			return acc
		}
		if level >= len(acc) {
			acc = append(acc, []int{})
		}
		acc[level] = append(acc[level], root.Val)
		acc = loop(root.Left, level+1, acc)
		acc = loop(root.Right, level+1, acc)
		return acc
	}

	ans := loop(root, 0, [][]int{})
	idx := 1
	for idx < len(ans) {
		nodes := ans[idx]
		for i := 0; i < len(nodes)/2; i++ {
			nodes[i], nodes[len(nodes)-i-1] = nodes[len(nodes)-i-1], nodes[i]
		}
		idx += 2
	}

	return ans
}
