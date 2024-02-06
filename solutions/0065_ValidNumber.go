package solutions

import "fmt"

func IsNumber(s string) bool {
	// sign part
	signIdx := 0
	if signIdx < len(s) && (s[signIdx] == '+' || s[signIdx] == '-') {
		signIdx++
	}

	// integer part
	integerIdx := signIdx
	for integerIdx < len(s) && s[integerIdx] >= '0' && s[integerIdx] <= '9' {
		integerIdx++
	}
	integerExisted := integerIdx != signIdx

	// dot part
	dotIdx := integerIdx
	if dotIdx < len(s) && s[dotIdx] == '.' {
		dotIdx++
	}
	dotExisted := dotIdx != integerIdx

	// decimal part
	decimalIdx := dotIdx
	for decimalIdx < len(s) && s[decimalIdx] >= '0' && s[decimalIdx] <= '9' {
		decimalIdx++
	}
	decimalExisted := decimalIdx != dotIdx
	if !integerExisted && !(dotExisted && decimalExisted) {
		fmt.Printf("%v, %v, %v\n", integerExisted, dotExisted, decimalExisted)
		return false
	}

	// e part
	eIdx := decimalIdx
	if eIdx < len(s) && (s[eIdx] == 'e' || s[eIdx] == 'E') {
		eIdx++
	}
	eExisted := eIdx != decimalIdx

	// exponent sign part
	exponentSignIdx := eIdx
	if exponentSignIdx < len(s) && (s[exponentSignIdx] == '+' || s[exponentSignIdx] == '-') {
		exponentSignIdx++
	}
	exponentSignExisted := exponentSignIdx != eIdx

	// exponent part
	exponentIdx := exponentSignIdx
	for exponentIdx < len(s) && s[exponentIdx] <= '9' && s[exponentIdx] >= '0' {
		exponentIdx++
	}
	exponentExisted := exponentIdx != exponentSignIdx

	return exponentIdx == len(s) && ((eExisted && exponentExisted) || !(eExisted || exponentSignExisted || exponentExisted))
}
