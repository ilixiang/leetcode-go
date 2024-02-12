package solutions

func IsScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	n := len(s1)
	cache := make([][][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = make([]int, n+1)
		}
	}

	var dp func(int, int, int) bool
	dp = func(i1 int, i2 int, length int) bool {
		if length == 1 {
			return s1[i1] == s2[i2]
		}

		cached := cache[i1][i2][length]
		if cached != 0 {
			return cached == 1
		}

		l := 1
		for l < length &&
			(!dp(i1, i2, l) || !dp(i1+l, i2+l, length-l)) &&
			(!dp(i1, i2+length-l, l) || !dp(i1+l, i2, length-l)) {
			l++
		}
		if l == length {
			cache[i1][i2][length] = -1
			return false
		} else {
			cache[i1][i2][length] = 1
			return true
		}
	}
	return dp(0, 0, len(s1))
}
