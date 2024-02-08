package solutions

func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool
	dfs = func(row int, col int, idx int) bool {
		if used[row][col] || word[idx] != board[row][col] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}

		used[row][col] = true
		rev := (row > 0 && dfs(row-1, col, idx+1)) ||
			(row < m-1 && dfs(row+1, col, idx+1)) ||
			(col > 0 && dfs(row, col-1, idx+1)) ||
			(col < n-1 && dfs(row, col+1, idx+1))
		used[row][col] = false
		return rev
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}
