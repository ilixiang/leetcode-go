package solutions

func CombinationSumIII(k int, n int) [][]int {
	res := make([][]int, 0)
	buf := make([]int, k)

	var dfs func(int, int, int)
	dfs = func(idx int, start int, target int) {
		if idx == k && target == 0 {
			tmp := make([]int, k)
			copy(tmp, buf)
			res = append(res, tmp)
			return
		}

		if target < start || idx == k {
			return
		}

		end := 10 - k + idx
		for i := start; i <= end; i++ {
			buf[idx] = i
			dfs(idx+1, i+1, target-i)
		}
	}
	dfs(0, 1, n)
	return res
}
