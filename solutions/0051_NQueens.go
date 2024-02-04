package solutions

func SolveNQueens(n int) [][]string {
	board := make([]int, n)
	results := make([][]string, 0)

	var recursive func(int)
	recursive = func(row int) {
		if row == n {
			result := make([]string, 0, n)
			buf := make([]byte, n)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if board[i]&(1<<j) == 0 {
						buf[j] = '.'
					} else {
						buf[j] = 'Q'
					}
				}
				result = append(result, string(buf))
			}
			results = append(results, result)
		}
		for col := 0; col < n; col++ {
			mask := 1 << col
			i := 0
			for i < row &&
				mask&board[i] == 0 &&
				mask&(board[i]<<(row-i)) == 0 &&
				mask&(board[i]>>(row-i)) == 0 {
				i++
			}
			if i == row {
				board[row] = mask
				recursive(row + 1)
				board[row] = 0
			}
		}
	}

	recursive(0)
	return results
}
