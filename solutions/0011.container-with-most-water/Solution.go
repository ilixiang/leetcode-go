package solutions

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	rev := 0
	for left < right {
		rev = max(rev, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return rev
}
