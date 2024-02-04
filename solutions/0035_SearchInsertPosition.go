package solutions

func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	pos := len(nums)
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid - 1
			pos = mid
		} else {
			left = mid + 1
		}
	}
	return pos
}
