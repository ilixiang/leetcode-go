package solutions

func IsMatchII(s string, p string) bool {
	cache := make([][]int, len(s)+1)
	for i := range cache {
		cache[i] = make([]int, len(p)+1)
	}

	var match func(int, int) bool
	match = func(i, j int) bool {
		if cache[i][j] != 0 {
			return cache[i][j] == 1
		}

		var rev bool
		if i == len(s) && j == len(p) {
			rev = true
		} else if j == len(p) {
			rev = false
		} else if i == len(s) && p[j] == '*' {
			rev = match(i, j+1)
		} else if i == len(s) {
			rev = false
		} else if s[i] == p[j] || p[j] == '?' {
			rev = match(i+1, j+1)
		} else if p[j] == '*' {
			rev = match(i+1, j) || match(i, j+1) || match(i+1, j+1)
		} else {
			rev = false
		}

		if rev {
			cache[i][j] = 1
		} else {
			cache[i][j] = -1
		}
		return rev
	}
	return match(0, 0)
}
