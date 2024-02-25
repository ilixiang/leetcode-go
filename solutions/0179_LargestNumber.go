package solutions

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	numstrs := make([]string, len(nums))
	for i := range numstrs {
		numstrs[i] = strconv.Itoa(nums[i])
	}

	compare := func(i1 int, i2 int) bool {
		return numstrs[i1]+numstrs[i2] > numstrs[i2]+numstrs[i1]
	}
	sort.Slice(numstrs, compare)
	joined := strings.Join(numstrs, "")
	idx := 0
	for idx < len(joined)-1 && joined[idx] == '0' {
		idx++
	}
	return joined[idx:]
}
