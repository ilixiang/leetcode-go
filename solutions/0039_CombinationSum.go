package solutions

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	buf := []int{}
	sort.Ints(candidates)
	var recursive func(int, int)
	recursive = func(s int, start int) {
		if s == 0 {
			result := make([]int, len(buf))
			copy(result, buf)
			results = append(results, result)
			return
		}

		if s < 0 {
			return
		}

		for i := start; i < len(candidates) && candidates[i] <= s; i++ {
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}
			buf = append(buf, candidates[i])
			recursive(s-candidates[i], i)
			buf = buf[:len(buf)-1]
		}
	}

	recursive(target, 0)
	return results
}
