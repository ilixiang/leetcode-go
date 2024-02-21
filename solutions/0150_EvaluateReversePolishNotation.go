package solutions

import "strconv"

func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "+" {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], b+a)
		} else if token == "-" {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], b-a)
		} else if token == "*" {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], b*a)
		} else if token == "/" {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], b/a)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
