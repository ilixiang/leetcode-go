package solutions

func Reverse(x int) int {
	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	rev := 0
	const MaxIntRemainder = int(MaxInt % 10)
	const MaxIntQuotient = int(MaxInt / 10)
	const MinIntRemainder = int(MinInt % 10)
	const MinIntQuotient = int(MinInt / 10)
	if x > 0 {
		for x != 0 {
			quotient := x / 10
			remainder := x % 10
			if rev < MaxIntQuotient || (rev == MaxIntQuotient && remainder <= MaxIntRemainder) {
				rev = rev*10 + remainder
			} else {
				return 0
			}
			x = quotient
		}
	} else if x < 0 {
		for x != 0 {
			quotient := x / 10
			remainder := x % 10
			if rev > MinIntQuotient || (rev == MinIntQuotient && remainder >= MinIntQuotient) {
				rev = rev*10 + remainder
			} else {
				return 0
			}
			x = quotient
		}
	}
	return rev
}
