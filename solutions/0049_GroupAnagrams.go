package solutions

import "sort"

func GroupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		sorted := string(bs)
		m[sorted] = append(m[sorted], s)
	}

	results := make([][]string, 0, len(m))
	for _, val := range m {
		results = append(results, val)
	}
	return results
}
