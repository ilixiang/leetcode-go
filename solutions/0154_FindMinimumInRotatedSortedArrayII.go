package solutions

func FindMinII(nums []int) int {
	res := int(uint(1<<31) - 1)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[left] <= nums[mid] && nums[mid] <= nums[right] {
			res = min(res, nums[left])
			if nums[left] != nums[right] {
				return res
			}
			for left <= right && nums[left] == nums[right] {
				left, right = left+1, right-1
			}
		} else if nums[left] > nums[mid] {
			res = min(res, nums[mid])
			right = mid - 1
		} else {
			res = min(res, nums[right])
			left, right = mid+1, right-1
		}
	}
	return res
}
