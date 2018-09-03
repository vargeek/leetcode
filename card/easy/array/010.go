package easy

// 有效的数独
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/30/
// https://leetcode-cn.com/problems/valid-sudoku/description/

func isValidSudoku(board [][]byte) bool {
	rows := [9][10]bool{}
	cols := [9][10]bool{}
	grids := [9][10]bool{}

	for row := 0; row < 9; row++ {
		r := row / 3
		for col := 0; col < 9; col++ {
			char := board[row][col]
			if char < '0' {
				continue
			}
			val := char - '0'

			if rows[row][val] {
				return false
			}
			rows[row][val] = true

			if cols[col][val] {
				return false
			}
			cols[col][val] = true

			c := col / 3
			grid := r*3 + c
			if grids[grid][val] {
				return false
			}
			grids[grid][val] = true
		}
	}

	return true
}
