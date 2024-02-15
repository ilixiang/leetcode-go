package solutions

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	s := map[int]bool{}
	for _, num := range nums {
		s[num] = true
	}

	res := 1
	for _, num := range nums {
		left, right := num-1, num+1
		for ; s[left]; left-- {
			s[left] = false
		}
		for ; s[right]; right++ {
			s[right] = false
		}
		res = max(res, right-left-1)
	}
	return res
}
