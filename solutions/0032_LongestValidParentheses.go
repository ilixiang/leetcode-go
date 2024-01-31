package solutions

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	stack := make([]int, 0, len(s))
	memo := make([]int, len(s)+1)

	for i := range s {
		c := s[i]
		if c == '(' {
			stack = append(stack, i)
			memo[i+1] = 0
		} else {
			if len(stack) != 0 {
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				memo[i+1] = i - left + 1 + memo[left]
			} else {
				memo[i+1] = 0
			}
		}
	}

	rev := 0
	for i := range memo {
		if memo[i] > rev {
			rev = memo[i]
		}
	}
	return rev
}
