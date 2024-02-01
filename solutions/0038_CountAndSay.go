package solutions

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := "1"
	for i := 1; i < n; i++ {
		sb := strings.Builder{}
		count := 0
		for j := 0; j < len(prev); j++ {
			if j == 0 || prev[j] == prev[j-1] {
				count++
			} else {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteByte(prev[j-1])
				count = 1
			}
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(prev[len(prev)-1])
		prev = sb.String()
	}
	return prev
}
