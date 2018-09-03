package easy

// 旋转图像
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/31/
// https://leetcode-cn.com/problems/rotate-image/

func rotateImage1(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < (length+1)/2; i++ {
		for j := 0; j < length/2; j++ {
			bak := matrix[i][j]
			ii := i
			jj := j
			for k := 0; k < 3; k++ {
				nextI := length - jj - 1
				nextJ := ii
				matrix[ii][jj] = matrix[nextI][nextJ]
				ii, jj = nextI, nextJ
			}
			matrix[ii][jj] = bak
		}
	}
}

func rotateImage(matrix [][]int) {
	length := len(matrix)

	l1 := (length + 1) / 2
	l2 := length / 2
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			bak := matrix[i][j]
			// i, j <- length-j-1, i
			matrix[i][j] = matrix[length-j-1][i]
			matrix[length-j-1][i] = matrix[length-i-1][length-j-1]
			matrix[length-i-1][length-j-1] = matrix[j][length-i-1]
			matrix[j][length-i-1] = bak
		}
	}
}
