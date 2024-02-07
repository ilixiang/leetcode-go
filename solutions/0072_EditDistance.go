package solutions

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cache := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		cache[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
		cache[i][n] = m - i
	}
	for i := 0; i < n; i++ {
		cache[m][i] = n - i
	}
	cache[m][n] = 0

	var dp func(int, int) int
	dp = func(i1 int, i2 int) int {
		cached := cache[i1][i2]
		if cached != -1 {
			return cached
		}

		if word1[i1] == word2[i2] {
			cached = dp(i1+1, i2+1)
		} else {
			insertionStep := dp(i1, i2+1)
			deletionStep := dp(i1+1, i2)
			replacementStep := dp(i1+1, i2+1)
			cached = 1 + min(insertionStep, deletionStep, replacementStep)
		}
		cache[i1][i2] = cached
		return cached
	}
	return dp(0, 0)
}
