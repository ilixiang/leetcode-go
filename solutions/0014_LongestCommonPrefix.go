package solutions

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sample := strs[0]
	sb := strings.Builder{}
	for i := range sample {
		cur := sample[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != cur {
				return sb.String()
			}
		}
		sb.WriteByte(cur)
	}
	return sb.String()
}
