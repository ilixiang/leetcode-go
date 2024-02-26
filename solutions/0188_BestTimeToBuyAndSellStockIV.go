package solutions

func MaxProfitIV(k int, prices []int) int {
	m := len(prices)
	cache := make([][][2]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([][2]int, k)
	}

	for i := 0; i < k; i++ {
		cache[m-1][i] = [2]int{0, prices[m-1]}
	}

	for i := m - 2; i >= 0; i-- {
		nextIdx := cache[i+1][k-1]
		cache[i][k-1] = [2]int{max(-prices[i]+nextIdx[1], nextIdx[0]), max(prices[i], nextIdx[1])}
		for j := k - 2; j >= 0; j-- {
			nextTx := cache[i+1][j+1]
			nextIdx := cache[i+1][j]
			cache[i][j] = [2]int{max(-prices[i]+nextIdx[1], nextIdx[0]), max(prices[i]+nextTx[0], nextIdx[1])}
		}
	}
	return cache[0][0][0]
}
