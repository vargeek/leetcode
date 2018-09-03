package others

// Pascal's Triangle
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/601/

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	result := [][]int{
		[]int{1},
	}

	for i := 1; i < numRows; i++ {
		result = append(result, []int{1})
		for j := 1; j < i; j++ {
			result[i] = append(result[i], result[i-1][j]+result[i-1][j-1])
		}
		result[i] = append(result[i], 1)
	}

	return result
}
