package treesandgraphs

// Number of Islands
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/792/

func numIslands(grid [][]byte) int {

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '1' {
				grid[row][col] = '0'
				count++
				stack := []int{row, col}
				idx := 0
				for idx < len(stack) {
					r := stack[idx]
					c := stack[idx+1]
					idx += 2

					if r-1 >= 0 && grid[r-1][c] == '1' {
						grid[r-1][c] = 0
						stack = append(stack, r-1, c)
					}
					if r+1 < len(grid) && grid[r+1][c] == '1' {
						grid[r+1][c] = 0
						stack = append(stack, r+1, c)
					}
					if c-1 >= 0 && grid[r][c-1] == '1' {
						grid[r][c-1] = 0
						stack = append(stack, r, c-1)
					}
					if c+1 < len(grid[r]) && grid[r][c+1] == '1' {
						grid[r][c+1] = 0
						stack = append(stack, r, c+1)
					}

				}
			}
		}
	}
	return count
}
