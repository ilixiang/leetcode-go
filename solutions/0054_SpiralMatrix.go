package solutions

func SpiralOrder(matrix [][]int) []int {
	height := len(matrix)
	width := len(matrix[0])

	rev := make([]int, 0, height*width)

	left, right := 0, width-1
	top, bottom := 0, height-1
	for top < bottom && left < right {
		for col := left; col < right; col++ {
			rev = append(rev, matrix[top][col])
		}
		for row := top; row < bottom; row++ {
			rev = append(rev, matrix[row][right])
		}
		for col := right; col > left; col-- {
			rev = append(rev, matrix[bottom][col])
		}
		for row := bottom; row > top; row-- {
			rev = append(rev, matrix[row][left])
		}
		left, right = left+1, right-1
		top, bottom = top+1, bottom-1
	}

	if left == right {
		for row := top; row <= bottom; row++ {
			rev = append(rev, matrix[row][left])
		}
	} else if top == bottom {
		for col := left; col <= right; col++ {
			rev = append(rev, matrix[top][col])
		}
	}

	return rev
}
