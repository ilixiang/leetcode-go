package solutions

func SearchInRotatedSortedArrayII(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			return true
		} else if nums[left] == nums[right] {
			if nums[left] == target {
				return true
			} else {
				left++
				right--
			}
		} else if target < nums[mid] {
			if (nums[left] <= nums[mid] && nums[mid] <= nums[right]) || nums[left] > nums[mid] {
				right = mid - 1
			} else if nums[left] < target {
				right = mid - 1
			} else if nums[left] > target {
				left = mid + 1
			} else {
				return true
			}
		} else {
			if (nums[left] <= nums[mid] && nums[mid] <= nums[right]) || nums[mid] > nums[right] {
				left = mid + 1
			} else if nums[right] < target {
				right = mid - 1
			} else if nums[right] > target {
				left = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
