package solutions

func TwoSumII(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		s := nums[left] + nums[right]
		if s == target {
			return []int{left, right}
		} else if s > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
