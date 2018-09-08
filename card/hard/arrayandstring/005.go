package hard

// Game of Life
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/831/
// https://leetcode.com/problems/game-of-life/description/

func gameOfLife(board [][]int) {
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])

	neighbors := [][2]int{
		[2]int{-1, -1},
		[2]int{-1, 0},
		[2]int{-1, 1},
		[2]int{0, -1},
		[2]int{0, 1},
		[2]int{1, -1},
		[2]int{1, 0},
		[2]int{1, 1},
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeighbors := 0
			for _, delta := range neighbors {
				r := row + delta[0]
				c := col + delta[1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					liveNeighbors += board[r][c] & 0x01
				}
			}

			switch liveNeighbors {
			case 2:
				board[row][col] |= (board[row][col] & 0x01) << 1
			case 3:
				board[row][col] |= 0x02
			}
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col] >>= 1
		}
	}
}
