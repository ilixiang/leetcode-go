package solutions

func WordBreak(s string, wordDict []string) bool {
	lookup := map[string]struct{}{}
	m, n := 20, 1
	for _, word := range wordDict {
		lookup[word] = struct{}{}
		m = min(m, len(word))
		n = max(n, len(word))
	}

	cache := make([]int, len(s)+1)
	cache[len(s)] = 1

	var dp func(int) bool
	dp = func(start int) bool {
		cached := cache[start]
		if cached != 0 {
			return cached == 1
		}

		cached = -1
		length := m
		for ; length <= n && start+length <= len(s); length++ {
			word := s[start : start+length]
			_, existed := lookup[word]
			if existed && dp(start+length) {
				cached = 1
				break
			}
		}
		cache[start] = cached
		return cached == 1
	}
	return dp(0)
}
