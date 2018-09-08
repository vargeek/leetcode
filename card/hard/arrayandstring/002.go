package hard

// Spiral Matrix
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/828/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])
	output := make([]int, 0, rows*cols)
	left, right, top, bottom := 0, cols-1, 0, rows-1
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			output = append(output, matrix[top][col])
		}
		if bottom == top {
			break
		}
		for row := top + 1; row <= bottom; row++ {
			output = append(output, matrix[row][right])
		}
		if left == right {
			break
		}
		for col := right - 1; col >= left; col-- {
			output = append(output, matrix[bottom][col])
		}
		for row := bottom - 1; row >= top+1; row-- {
			output = append(output, matrix[row][left])
		}
		left++
		right--
		top++
		bottom--
	}

	return output
}
