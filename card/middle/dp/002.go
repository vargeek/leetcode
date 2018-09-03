package dp

// Unique Paths
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/808/

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	paths := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		paths = append(paths, make([]int, n))
	}
	paths[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				paths[i][j] = paths[i-1][j]
			}
			if j > 0 {
				paths[i][j] += paths[i][j-1]
			}
		}
	}

	return paths[m-1][n-1]
}
