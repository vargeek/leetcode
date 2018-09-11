package treesandgraphs

// Surrounded Regions
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/843/

func solve(board [][]byte) {

	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])
	if cols == 0 {
		return
	}

	dir := []int{0, -1, -1, 0, 0, 1, 1, 0}
	fill := func(row, col int) {
		stack := []int{row, col}
		board[row][col] = ' '
		for len(stack) > 0 {
			r, c := stack[0], stack[1]
			for i := 0; i < len(dir); i += 2 {
				rr, cc := r+dir[i], c+dir[i+1]
				if rr >= 0 && rr < rows && cc >= 0 && cc < cols {
					aByte := board[rr][cc]
					if aByte == 'O' || aByte == '*' {
						board[rr][cc] = ' '
						stack = append(stack, rr, cc)
					}
				}
			}
			stack = stack[2:]
		}
	}

	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' {
			fill(0, c)
		}
		if board[rows-1][c] == 'O' {
			fill(rows-1, c)
		}
	}
	for r := 0; r < rows; r++ {
		if board[r][0] == 'O' {
			fill(r, 0)
		}
		if board[r][cols-1] == 'O' {
			fill(r, cols-1)
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == ' ' {
				board[row][col] = 'O'
			}
		}
	}
}
