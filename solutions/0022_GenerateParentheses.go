package solutions

func generateParenthesis(n int) []string {
	rev := make([]string, 0)
	buf := make([]byte, 0, n*2)

	var generate func(int, int)
	generate = func(left int, right int) {
		if right == 0 {
			rev = append(rev, string(buf))
		} else if left >= right {
			buf = append(buf, '(')
			generate(left-1, right)
			buf = buf[:len(buf)-1]
		} else {
			if left != 0 {
				buf = append(buf, '(')
				generate(left-1, right)
				buf = buf[:len(buf)-1]
			}
			buf = append(buf, ')')
			generate(left, right-1)
			buf = buf[:len(buf)-1]
		}
	}
	generate(n, n)
	return rev
}
