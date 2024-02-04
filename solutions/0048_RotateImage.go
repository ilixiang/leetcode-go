package solutions

func Rotate(matrix [][]int) {
	size := len(matrix)

	low, high := 0, size-1
	for low < high {
		for col := 0; col < size; col++ {
			matrix[low][col], matrix[high][col] = matrix[high][col], matrix[low][col]
		}
		low++
		high--
	}

	for row := 0; row < size; row++ {
		for col := row + 1; col < size; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
}
