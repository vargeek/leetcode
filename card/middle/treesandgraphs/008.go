package treesandgraphs

// Longest Increasing Path in a Matrix
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/849/

func longestIncreasingPath1(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	lengths := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		lengths = append(lengths, make([]int, n))
	}
	longest := 0

	dir := []int{0, -1, -1, 0, 0, 1, 1, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if lengths[i][j] > 0 {
				continue
			}
			lengths[i][j] = 1
			stack := []int{i, j}
			for len(stack) > 0 {
				ii, jj := stack[0], stack[1]
				if lengths[ii][jj] > longest {
					longest = lengths[ii][jj]
				}
				stack = stack[2:]
				for di := 0; di < len(dir); di += 2 {
					iii := ii + dir[di]
					jjj := jj + dir[di+1]
					if iii >= 0 && iii < m && jjj >= 0 && jjj < n && matrix[iii][jjj] > matrix[ii][jj] && lengths[ii][jj]+1 > lengths[iii][jjj] {
						lengths[iii][jjj] = lengths[ii][jj] + 1
						stack = append(stack, iii, jjj)
					}
				}
			}
		}
	}

	return longest
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	lengths := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		lengths = append(lengths, make([]int, n))
	}

	dir := []int{0, -1, -1, 0, 0, 1, 1, 0}
	var dp func(int, int) int
	dp = func(row, col int) int {
		length := lengths[row][col]
		if length > 0 {
			return length
		}

		val := matrix[row][col]
		for di := 0; di < len(dir); di += 2 {
			r := row + dir[di]
			c := col + dir[di+1]
			if r >= 0 && r < m && c >= 0 && c < n && matrix[r][c] < val {
				l := dp(r, c)
				if l > length {
					length = l
				}
			}
		}
		lengths[row][col] = length + 1
		return length + 1
	}

	longest := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			length := dp(row, col)
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
