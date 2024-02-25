package solutions

import (
	"strconv"
	"strings"
)

func FractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	positive := (numerator >= 0 && denominator > 0) || (numerator <= 0 && denominator < 0)
	if numerator > 0 {
		numerator = -numerator
	}
	if denominator > 0 {
		denominator = -denominator
	}
	integerPart := strconv.Itoa(numerator / denominator)
	numerator = (numerator % denominator) * 10

	numeratorMap := map[int]int{}
	stack := []string{}

	existed := false
	for numerator != 0 && !existed {
		numeratorMap[numerator] = len(stack)
		stack = append(stack, strconv.Itoa(numerator/denominator))
		numerator = (numerator % denominator) * 10
		_, existed = numeratorMap[numerator]
	}

	prefix := ""
	if !positive {
		prefix = "-"
	}
	if numerator == 0 {
		if len(stack) == 0 {
			return prefix + integerPart
		} else {
			return prefix + integerPart + "." + strings.Join(stack, "")
		}
	} else {
		idx := numeratorMap[numerator]
		return prefix + integerPart + "." + strings.Join(stack[:idx], "") + "(" + strings.Join(stack[idx:], "") + ")"
	}
}
