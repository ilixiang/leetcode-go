package solutions

func NumTrees(n int) int {
	cache := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = -1
	}
	cache[0] = 1
	cache[1] = 1

	var count func(int) int
	count = func(x int) int {
		cached := cache[x]
		if cached != -1 {
			return cached
		}

		num := 0
		for i := 0; i <= x-1; i++ {
			num += count(i) * count(x-i-1)
		}
		cache[x] = num
		return num
	}
	return count(n)
}
