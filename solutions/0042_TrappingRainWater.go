package solutions

func Trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	water := 0
	left, right := 0, len(height)-1
	maxLeftHeight, maxRightHeight := height[left], height[right]
	for left < right {
		if height[left] < maxLeftHeight {
			water += maxLeftHeight - height[left]
		} else {
			maxLeftHeight = height[left]
		}
		if height[right] < maxRightHeight {
			water += maxRightHeight - height[right]
		} else {
			maxRightHeight = height[right]
		}

		if maxLeftHeight < maxRightHeight {
			left++
		} else {
			right--
		}
	}
	return water
}
