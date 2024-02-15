package solutions

func MinCut(s string) int {
	m := len(s)

	parlindromeCache := make([][]int, m)
	for i := 0; i < m; i++ {
		parlindromeCache[i] = make([]int, m)
	}

	var isParlindrome func(int, int) bool
	isParlindrome = func(left int, right int) bool {
		cached := parlindromeCache[left][right]
		if cached != 0 {
			return cached == 1
		}

		cached = -1
		if left >= right || (s[left] == s[right] && isParlindrome(left+1, right-1)) {
			cached = 1
		}
		parlindromeCache[left][right] = cached
		return cached == 1
	}

	candidates := make([][]int, m)
	for left := 0; left < m; left++ {
		for right := left; right < m; right++ {
			if isParlindrome(left, right) {
				candidates[left] = append(candidates[left], right)
			}
		}
	}

	cache := make([]int, m+1)
	for i := 0; i < m; i++ {
		cache[i] = -1
	}
	cache[m] = 0
	var dp func(int) int
	dp = func(i int) int {
		cached := cache[i]
		if cached != -1 {
			return cached
		}

		cached = m
		for _, right := range candidates[i] {
			cached = min(cached, 1+dp(right+1))
		}
		cache[i] = cached
		return cached
	}
	return dp(0) - 1
}
