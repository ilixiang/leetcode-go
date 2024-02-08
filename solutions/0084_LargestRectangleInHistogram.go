package solutions

func LargestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights))
	i := 0
	result := 0
	for i < len(heights) {
		height := heights[i]
		if len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) != 0 {
				left = stack[len(stack)-1]
			}
			area := (i - left - 1) * heights[idx]
			if result < area {
				result = area
			}
		} else {
			stack = append(stack, i)
			i++
		}
	}

	left, right := -1, len(heights)
	for i := 0; i < len(stack); i++ {
		area := (right - left - 1) * heights[stack[i]]
		if result < area {
			result = area
		}
		left = stack[i]
	}
	return result
}
