package treesandgraphs

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const null = 1 << 31

func newTreeNodeFromInts(nums ...int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	stack := []*TreeNode{&TreeNode{nums[0], nil, nil}}
	stackI := 0
	numsI := 1
	for stackI < len(stack) && numsI < len(nums) {
		num := nums[numsI]
		if num != null {
			stack[stackI].Left = &TreeNode{num, nil, nil}
			stack = append(stack, stack[stackI].Left)
		}
		numsI++
		if numsI < len(nums) {
			num = nums[numsI]
			if num != null {
				stack[stackI].Right = &TreeNode{num, nil, nil}
				stack = append(stack, stack[stackI].Right)
			}
			numsI++
		}
		stackI++
	}

	return stack[0]
}
