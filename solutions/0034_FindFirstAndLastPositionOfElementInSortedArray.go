package solutions

func searchRange(nums []int, target int) []int {

	left, right := 0, len(nums)-1
	divider := -1
	for left <= right {
		divider = (left + right) >> 1
		if nums[divider] == target {
			break
		} else if nums[divider] < target {
			left = divider + 1
		} else {
			right = divider - 1
		}
	}

	if left > right {
		return []int{-1, -1}
	}

	first, last := -1, -1
	low, high := left, divider
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			first = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	low, high = divider, right
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			last = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return []int{first, last}
}
