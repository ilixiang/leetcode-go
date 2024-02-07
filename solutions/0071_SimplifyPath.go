package solutions

import "strings"

func SimplifyPath(path string) string {
	stack := make([]string, 0, 3000)
	i, j := 0, 0
	for j <= len(path) {
		if j != len(path) && path[j] != '/' {
			j++
		} else {
			part := path[i:j]
			if part == ".." {
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			} else if len(part) != 0 && part != "." {
				stack = append(stack, part)
			}
			i = j + 1
			j = i
		}
	}
	return "/" + strings.Join(stack, "/")
}
