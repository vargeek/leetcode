package arrayandstring

// Set Matrix Zeroes

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/777/
// https://leetcode.com/problems/set-matrix-zeroes/description/

func visitMatrix(matrix [][]int, row int, col int) {
	for row < len(matrix) && col < len(matrix[row]) {
		if matrix[row][col] == 0 {
			break
		}
		if col+1 < len(matrix[row]) {
			col++
		} else {
			row++
			col = 0
		}
	}

	if row < len(matrix) && col < len(matrix[row]) {
		if col+1 < len(matrix[row]) {
			visitMatrix(matrix, row, col+1)
		} else {
			visitMatrix(matrix, row+1, 0)
		}
		setZerosToRowAndColumn(matrix, row, col)
	}
}
func setZerosToRowAndColumn(matrix [][]int, row int, col int) {
	for i := 0; i < len(matrix[row]); i++ {
		matrix[row][i] = 0
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}

func setZeroes(matrix [][]int) {
	visitMatrix(matrix, 0, 0)
}
