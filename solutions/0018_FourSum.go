package solutions

import "sort"

func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return make([][]int, 0)
	}

	rev := make([][]int, 0)
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++ {
		if a != 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b != a+1 && nums[b] == nums[b-1] {
				continue
			}
			c, d := b+1, len(nums)-1
			for c < d {
				if c != b+1 && nums[c] == nums[c-1] {
					c++
					continue
				}
				if d != len(nums)-1 && nums[d] == nums[d+1] {
					d--
					continue
				}
				s := nums[a] + nums[b] + nums[c] + nums[d]
				if s == target {
					rev = append(rev, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
				} else if s < target {
					c++
				} else {
					d--
				}
			}
		}
	}
	return rev
}
