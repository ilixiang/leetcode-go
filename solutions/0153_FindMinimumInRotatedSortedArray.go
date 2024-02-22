package solutions

func FindMin(nums []int) int {
	res := int(uint(1<<31) - 1)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[left] <= nums[mid] && nums[mid] <= nums[right] {
			return min(res, nums[left])
		} else if nums[left] > nums[mid] {
			res = min(res, nums[mid])
			right = mid - 1
		} else {
			res = min(res, nums[mid])
			left = mid + 1
		}
	}
	return res
}
