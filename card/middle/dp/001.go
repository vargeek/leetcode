package dp

// Jump Game
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/807/
// https://leetcode.com/problems/jump-game/description/

func canJump1(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	reachable := make([]bool, len(nums))
	reachable[0] = true
	for i := 0; i < len(nums)-1; i++ {
		if reachable[i] {
			for j := 1; j <= nums[i]; j++ {
				if i+j < len(nums) {
					reachable[i+j] = true
				} else {
					return true
				}
			}
		}
	}

	return reachable[len(nums)-1]
}

func canJump(nums []int) bool {
	right := 0
	for i := 0; i < len(nums); i++ {
		if i > right {
			return false
		}
		if i+nums[i] > right {
			right = i + nums[i]
			if right >= len(nums)-1 {
				return true
			}
		}
	}
	return true
}
