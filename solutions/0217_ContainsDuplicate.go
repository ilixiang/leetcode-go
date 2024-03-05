package solutions

func ContainsDuplicate(nums []int) bool {
	s := map[int]struct{}{}
	for _, num := range nums {
		_, existed := s[num]
		if existed {
			return true
		}
		s[num] = struct{}{}
	}
	return false
}
