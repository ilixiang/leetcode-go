package solutions

func IsMatch(s string, p string) bool {
	cache := make([][]int, len(s)+1)
	for i := range cache {
		cache[i] = make([]int, len(p)+1)
	}

	var match func(int, int) bool
	match = func(i int, j int) bool {
		if i == len(s) && j == len(p) {
			return true
		}
		if i != len(s) && j == len(p) {
			return false
		}
		if cache[i][j] == 1 {
			return true
		} else if cache[i][j] == -1 {
			return false
		}

		var rev bool
		curMatched := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			rev = (curMatched && match(i+1, j)) || match(i, j+2)
		} else {
			rev = curMatched && match(i+1, j+1)
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
