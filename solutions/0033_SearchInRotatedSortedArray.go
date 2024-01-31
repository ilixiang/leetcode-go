package solutions

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[left] <= nums[mid] && nums[mid] <= nums[right] {
				right = mid - 1
			} else if nums[left] > nums[mid] {
				right = mid - 1
			} else {
				if target < nums[right] {
					left = mid + 1
				} else if target == nums[right] {
					return right
				} else {
					right = mid - 1
				}
			}
		} else {
			if nums[left] <= nums[mid] && nums[mid] <= nums[right] {
				left = mid + 1
			} else if nums[left] > nums[mid] {
				if nums[left] > target {
					left = mid + 1
				} else if nums[left] == target {
					return left
				} else {
					right = mid - 1
				}
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
