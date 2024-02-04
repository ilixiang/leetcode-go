package solutions

func Divide(dividend int, divisor int) int {
	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	if dividend == 0 {
		return 0
	}

	positive := true
	if dividend > 0 {
		dividend = -dividend
	} else {
		positive = false
	}
	if divisor > 0 {
		divisor = -divisor
	} else {
		if positive {
			positive = false
		} else {
			positive = true
		}
	}

	rev := 0
	for dividend <= divisor {
		leftShift := 0
		for (divisor << leftShift) >= dividend {
			leftShift++
		}
		rev -= 1 << (leftShift - 1)
		dividend -= divisor << (leftShift - 1)
	}
	if positive && rev != int(MinInt) {
		return 0 - rev
	} else if positive {
		return int(MaxInt)
	} else {
		return rev
	}
}
