package solutions

func IsInterleave(s1 string, s2 string, s3 string) bool {
	m, n, o := len(s1), len(s2), len(s3)
	if m+n != o {
		return false
	}

	cache := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		cache[i] = make([]int, n+1)
	}
	cache[m][n] = 1

	var compare func(int, int) int
	compare = func(i1 int, i2 int) int {
		cached := cache[i1][i2]
		if cached != 0 {
			return cached
		}

		i3 := i1 + i2
		if i1 == m && (i2 == n || s2[i2:] == s3[i3:]) {
			cached = 1
		} else if i2 == n && (i1 == m || s1[i1:] == s3[i3:]) {
			cached = 1
		} else if i1 == m || i2 == n {
			cached = -1
		} else if (s1[i1] == s3[i3] && compare(i1+1, i2) == 1) ||
			(s2[i2] == s3[i3] && compare(i1, i2+1) == 1) {
			cached = 1
		} else {
			cached = -1
		}

		cache[i1][i2] = cached
		return cached
	}
	return compare(0, 0) == 1
}
