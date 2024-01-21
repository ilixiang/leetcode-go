package solutions

import "unicode"

func myAtoi(s string) int {
	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	rev := 0
	const MaxIntRemainder = int(MaxInt % 10)
	const MaxIntQuotient = int(MaxInt / 10)
	const MinIntRemainder = int(MinInt % 10)
	const MinIntQuotient = int(MinInt / 10)

	rs := []rune(s)
	positive := true

	i := 0
	// ignore leading space
	for i < len(rs) && unicode.IsSpace(rs[i]) {
		i++
	}

	// check + and -
	if i < len(s) {
		r := rs[i]
		if r == '-' {
			positive = false
			i++
		} else if r == '+' {
			i++
		}
	}

	// read until non-digit
	if positive {
		for ; i < len(s) && unicode.IsDigit(rs[i]); i++ {
			digit := int(rs[i] - '0')
			if rev < MaxIntQuotient || (rev == MaxIntQuotient && digit <= MaxIntRemainder) {
				rev = rev*10 + digit
			} else {
				return int(MaxInt)
			}
		}
	} else {
		for ; i < len(s) && unicode.IsDigit(rs[i]); i++ {
			digit := int(rs[i] - '0')
			if rev > MinIntQuotient || (rev == MinIntQuotient && -1*digit >= MinIntRemainder) {
				rev = rev*10 - digit
			} else {
				return int(MinInt)
			}
		}
	}
	return rev
}
