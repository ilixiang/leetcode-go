package solutions

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	rev := 0
	for left < right {
		leftHeight, rightHeight := height[left], height[right]
		rev = max(rev, min(leftHeight, rightHeight)*(right-left))
		if leftHeight < rightHeight {
			left++
		} else {
			right--
		}
	}
	return rev
}
