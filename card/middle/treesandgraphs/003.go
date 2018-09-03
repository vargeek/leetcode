package treesandgraphs

// Construct Binary Tree from Preorder and Inorder Traversal

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/788/
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

func buildTree(preorder []int, inorder []int) *TreeNode {
	var loop func(int, int, int, int) *TreeNode
	loop = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preRight < preLeft {
			return nil
		}
		idx := inLeft
		for idx <= inRight {
			if inorder[idx] == preorder[preLeft] {
				break
			}
			idx++
		}
		leftCount := idx - inLeft
		leftTree := loop(preLeft+1, preLeft+leftCount, inLeft, idx-1)
		rightTree := loop(preLeft+leftCount+1, preRight, idx+1, inRight)

		return &TreeNode{Val: preorder[preLeft], Left: leftTree, Right: rightTree}
	}
	return loop(0, len(preorder)-1, 0, len(inorder)-1)
}
