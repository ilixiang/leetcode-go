package solutions

func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return (left + right) >> 1
}
