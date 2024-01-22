package solutions

func IsMatch(s string, p string) bool {
	var match func(int, int) bool
	match = func(i int, j int) bool {
		if i == len(s) && j == len(p) {
			return true
		}
		if i != len(s) && j == len(p) {
			return false
		}

		curMatched := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			return (curMatched && match(i+1, j)) || match(i, j+2)
		} else {
			return curMatched && match(i+1, j+1)
		}
	}
	return match(0, 0)
}
