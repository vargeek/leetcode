package backtracking

// Word Search II
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/119/backtracking/853/

func findWords1(board [][]byte, words []string) []string {

	rows := len(board)
	if rows == 0 {
		return []string{}
	}
	cols := len(board[0])

	result := make([]string, 0, len(words))
	var loop func(string, int, int, int) bool
	loop = func(word string, step int, row int, col int) bool {
		if word[step] != board[row][col] {
			return false
		}
		if step+1 >= len(word) {
			return true
		}

		board[row][col] -= 32
		valid := false
		if row > 0 {
			if loop(word, step+1, row-1, col) {
				valid = true
			}
		}
		if !valid && row+1 < rows {
			if loop(word, step+1, row+1, col) {
				valid = true
			}
		}
		if !valid && col > 0 {
			if loop(word, step+1, row, col-1) {
				valid = true
			}
		}
		if !valid && col+1 < cols {
			if loop(word, step+1, row, col+1) {
				valid = true
			}
		}
		board[row][col] += 32
		return valid
	}

	duplicate := map[string]bool{}
	for _, word := range words {
		if duplicate[word] {
			continue
		}
		duplicate[word] = true
	outer:
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				if loop(word, 0, row, col) {
					result = append(result, word)
					break outer
				}
			}
		}
	}
	return result
}

func findWords(board [][]byte, words []string) []string {
	result := make([]string, 0, len(words))
	rows := len(board)
	if rows == 0 {
		return result
	}
	cols := len(board[0])
	if cols == 0 {
		return result
	}

	type Tree struct {
		children [26]*Tree
		word     string
	}

	insert := func(tree *Tree, word string) {
		for i := 0; i < len(word); i++ {
			ch := word[i] - 'a'
			if tree.children[ch] == nil {
				tree.children[ch] = &Tree{}
			}
			tree = tree.children[ch]
		}
		tree.word = word
	}
	tree := &Tree{}
	for _, word := range words {
		insert(tree, word)
	}

	found := map[string]bool{}
	var search func(*Tree, int, int)
	search = func(tree *Tree, row, col int) {
		if row < 0 || row >= rows {
			return
		}
		if col < 0 || col >= cols {
			return
		}
		ch := board[row][col]
		if ch == 0 {
			return
		}
		tree = tree.children[ch-'a']
		if tree == nil {
			return
		}
		if tree.word != "" {
			found[tree.word] = true
		}
		board[row][col] = 0

		search(tree, row, col-1)
		search(tree, row, col+1)
		search(tree, row-1, col)
		search(tree, row+1, col)

		board[row][col] = ch
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			search(tree, row, col)
		}
	}
	for word := range found {
		result = append(result, word)
	}

	return result
}
