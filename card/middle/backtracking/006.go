package backtracking

// Word Search
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/797/
// https://leetcode.com/problems/word-search/

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	var check func(int, int, int) bool
	check = func(idx, row, col int) bool {
		if idx >= len(word) {
			return true
		}
		if row-1 >= 0 {
			if board[row-1][col] == word[idx] {
				board[row-1][col] = ' '
				success := check(idx+1, row-1, col)
				board[row-1][col] = word[idx]
				if success {
					return true
				}
			}
		}
		if row+1 < len(board) {
			if board[row+1][col] == word[idx] {
				board[row+1][col] = ' '
				success := check(idx+1, row+1, col)
				board[row+1][col] = word[idx]
				if success {
					return true
				}
			}
		}
		if col-1 >= 0 {
			if board[row][col-1] == word[idx] {
				board[row][col-1] = ' '
				success := check(idx+1, row, col-1)
				board[row][col-1] = word[idx]
				if success {
					return true
				}
			}
		}
		if col+1 < len(board[row]) {
			if board[row][col+1] == word[idx] {
				board[row][col+1] = ' '
				success := check(idx+1, row, col+1)
				board[row][col+1] = word[idx]
				if success {
					return true
				}
			}
		}

		return false
	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == word[0] {
				board[row][col] = ' '
				if check(1, row, col) {
					return true
				}
				board[row][col] = word[0]
			}
		}
	}
	return false
}
