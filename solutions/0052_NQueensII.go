package solutions

func TotalNQueens(n int) int {
	board := make([]int, n)

	var recursive func(int) int
	recursive = func(row int) int {
		if row == n {
			return 1
		}

		rev := 0
		for col := 0; col < n; col++ {
			mask := 1 << col
			i := 0
			for i < row &&
				board[i]&mask == 0 &&
				board[i]&(mask<<(row-i)) == 0 &&
				board[i]&(mask>>(row-i)) == 0 {
				i++
			}
			if i == row {
				board[row] = mask
				rev += recursive(row + 1)
				board[row] = 0
			}
		}
		return rev
	}

	return recursive(0)
}
