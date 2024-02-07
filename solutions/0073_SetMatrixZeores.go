package solutions

func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col := 0
	for col < n && matrix[0][col] != 0 {
		col++
	}
	firstRow := col < n

	row := 0
	for row < m && matrix[row][0] != 0 {
		row++
	}
	firstCol := row < m

	for row = 1; row < m; row++ {
		for col = 1; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[0][col] = 0
				matrix[row][0] = 0
			}
		}
	}

	for row = 1; row < m; row++ {
		if matrix[row][0] == 0 {
			for col = 1; col < n; col++ {
				matrix[row][col] = 0
			}
		}
	}
	for col = 1; col < n; col++ {
		if matrix[0][col] == 0 {
			for row = 1; row < m; row++ {
				matrix[row][col] = 0
			}
		}
	}

	if firstRow {
		for col = 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if firstCol {
		for row = 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}
