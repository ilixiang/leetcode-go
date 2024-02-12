package solutions

func RestoreIPAddress(s string) []string {
	m := len(s)

	END := []string{""}
	EMPTY := []string{}
	cache := make([][][]string, 4)
	for i := 0; i < 4; i++ {
		cache[i] = make([][]string, m+1)
	}

	var dp func(int, int) []string
	dp = func(idx int, pos int) []string {
		if pos == 4 && idx == m {
			return END
		}
		if pos == 4 || idx >= m || (4-pos)*3 < m-idx {
			return EMPTY
		}

		cached := cache[pos][idx]
		if cached != nil {
			return cached
		}

		cached = make([]string, 0, 4)

		prefix := "."
		if pos == 0 {
			prefix = ""
		}

		suffixes := dp(idx+1, pos+1)
		for _, suffix := range suffixes {
			cached = append(cached, prefix+s[idx:idx+1]+suffix)
		}

		if idx < m-1 && s[idx] != '0' {
			suffixes = dp(idx+2, pos+1)
			for _, suffix := range suffixes {
				cached = append(cached, prefix+s[idx:idx+2]+suffix)
			}
		}

		if idx < m-2 && (s[idx] == '1' || (s[idx] == '2' && (s[idx+1] < '5' || (s[idx+1] == '5' && s[idx+2] <= '5')))) {
			suffixes = dp(idx+3, pos+1)
			for _, suffix := range suffixes {
				cached = append(cached, prefix+s[idx:idx+3]+suffix)
			}
		}
		cache[pos][idx] = cached
		return cached
	}
	return dp(0, 0)
}
