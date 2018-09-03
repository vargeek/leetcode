package easytree

// Convert Sorted Array to Binary Search Tree

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/631/
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var loop func([]int, int, int) *TreeNode
	loop = func(nums []int, left int, right int) *TreeNode {
		if right-left+1 <= 0 {
			return nil
		}

		mid := (right-left+1)/2 + left
		return &TreeNode{
			Val:   nums[mid],
			Left:  loop(nums, left, mid-1),
			Right: loop(nums, mid+1, right),
		}
	}
	return loop(nums, 0, len(nums)-1)
}
