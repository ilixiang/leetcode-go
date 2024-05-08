package solutions

func ContainsNearbyDuplicate(nums []int, k int) bool {
	s := make(map[int]bool, k+1)
	for i := 0; i < k && i < len(nums); i++ {
		s[nums[i]] = true
	}
	if len(s) < k && len(s) < len(nums) {
		return true
	}

	for i := k; i < len(nums); i++ {
		if s[nums[i]] {
			return true
		}
		s[nums[i]] = true
		delete(s, nums[i-k])
	}
	return false
}
