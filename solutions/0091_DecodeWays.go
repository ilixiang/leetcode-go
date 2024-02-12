package solutions

func NumDecodings(s string) int {
	m := len(s)
	cache := make([]int, m+1)
	for i := 0; i < m; i++ {
		cache[i] = -1
	}
	cache[m] = 1

	var dp func(int) int
	dp = func(idx int) int {
		cached := cache[idx]
		if cached != -1 {
			return cached
		}

		c := s[idx]
		if c == '0' {
			cached = 0
		} else if idx < m-1 && (c == '1' || (c == '2' && int(s[idx+1]-'0') < 7)) {
			cached = dp(idx+1) + dp(idx+2)
		} else {
			cached = dp(idx + 1)
		}
		cache[idx] = cached
		return cached
	}
	return dp(0)
}
