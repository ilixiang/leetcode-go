package solutions

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	rev := make([][]int, 0)
	for l := 0; l < len(nums)-2; l++ {
		if l != 0 && nums[l] == nums[l-1] {
			continue
		}
		m, r := l+1, len(nums)-1
		for m < r {
			if m != l+1 && nums[m] == nums[m-1] {
				m++
				continue
			}
			if r != len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}

			sum := nums[l] + nums[m] + nums[r]
			if sum == 0 {
				rev = append(rev, []int{nums[l], nums[m], nums[r]})
				m++
				r--
			} else if sum < 0 {
				m++
			} else {
				r--
			}
		}
	}
	return rev
}
