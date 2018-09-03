package sortingandsearching

// Search a 2D Matrix II
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/806/
// https://leetcode.com/problems/search-a-2d-matrix-ii/description/

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	var loop func(int, int, int, int) bool
	loop = func(rowlo, collo, rowhi, colhi int) bool {
		if rowhi < rowlo || colhi < collo {
			return false
		}
		if target < matrix[rowlo][collo] || target > matrix[rowhi][colhi] {
			return false
		}
		rowmid, colmid := (rowlo+rowhi)/2, (collo+colhi)/2
		return matrix[rowmid][colmid] == target ||
			loop(rowlo, collo, rowmid-1, colmid) ||
			loop(rowlo, colmid+1, rowmid, colhi) ||
			loop(rowmid+1, colmid, rowhi, colhi) ||
			loop(rowmid, collo, rowhi, colmid-1)
	}
	return loop(0, 0, len(matrix)-1, len(matrix[0])-1)
}
