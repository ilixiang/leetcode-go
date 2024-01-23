package solutions

func RomanToInt(s string) int {
	romanIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		nums[i] = romanIntMap[s[i]]
	}

	rev := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		if nums[i] < prev {
			rev -= nums[i]
		} else {
			rev += nums[i]
		}
		prev = nums[i]
	}
	return rev
}
