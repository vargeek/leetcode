package treesandgraphs

// Friend Circles
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/846/

func findCircleNum(M [][]int) int {
	length := len(M)
	circles := make([]int, length)

	stack := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if circles[i] == 0 {
			circles[i] = i + 1
			stack = append(stack, i)
			for len(stack) > 0 {
				s := stack[0]
				stack = stack[1:]
				for j := 0; j < length; j++ {
					if M[s][j] == 1 && circles[j] == 0 {
						circles[j] = circles[s]
						stack = append(stack, j)
					}
				}
			}
		}
	}
	ids := make(map[int]bool, length)
	for _, circle := range circles {
		ids[circle] = true
	}
	return len(ids)
}
