package solutions

func MaximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	stack := make([]int, 0, n)
	result := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == '0' {
				heights[col] = 0
			} else {
				heights[col]++
			}
		}

		for col := 0; col < n; col++ {
			height := heights[col]
			for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
				poppedIdx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				left, right := -1, col
				if len(stack) > 0 {
					left = stack[len(stack)-1]
				}
				result = max(result, (right-left-1)*heights[poppedIdx])
			}
			stack = append(stack, col)
		}

		left, right := -1, n
		for i := 0; i < len(stack); i++ {
			height := heights[stack[i]]
			result = max(result, height*(right-left-1))
			left = stack[i]
		}
		stack = stack[:0]
	}
	return result
}
