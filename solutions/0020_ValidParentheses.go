package solutions

func IsValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		r := s[i]
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if r == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if r == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
