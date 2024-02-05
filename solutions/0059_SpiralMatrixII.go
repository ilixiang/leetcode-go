package solutions

func GenerateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	val := 1
	low, high := 0, n-1
	for low < high {
		for col := low; col < high; col++ {
			matrix[low][col] = val
			val++
		}
		for row := low; row < high; row++ {
			matrix[row][high] = val
			val++
		}
		for col := high; col > low; col-- {
			matrix[high][col] = val
			val++
		}
		for row := high; row > low; row-- {
			matrix[row][low] = val
			val++
		}
		low, high = low+1, high-1
	}

	if low == high {
		matrix[low][high] = val
	}

	return matrix
}
