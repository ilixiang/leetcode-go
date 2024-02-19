package solutions

import "strings"

func WordBreakII(s string, wordDict []string) []string {
	m, n := 10, 1
	lookup := map[string]struct{}{}
	for _, word := range wordDict {
		lookup[word] = struct{}{}
		m = min(m, len(word))
		n = max(n, len(word))
	}

	res := make([]string, 0)
	buf := make([]string, 0)
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			res = append(res, strings.Join(buf, " "))
			return
		}

		for length := m; length <= n && start+length <= len(s); length++ {
			word := s[start : start+length]
			_, existed := lookup[word]
			if existed {
				buf = append(buf, word)
				dfs(start + length)
				buf = buf[:len(buf)-1]
			}
		}
	}
	dfs(0)
	return res
}
