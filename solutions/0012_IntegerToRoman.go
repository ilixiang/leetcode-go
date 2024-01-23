package solutions

import "strings"

type RomanSymbol struct {
	Symbol string
	Value  int
}

func IntToRoman(num int) string {
	romanSymbols := [13]RomanSymbol{
		{Symbol: "M", Value: 1000},
		{Symbol: "CM", Value: 900},
		{Symbol: "D", Value: 500},
		{Symbol: "CD", Value: 400},
		{Symbol: "C", Value: 100},
		{Symbol: "XC", Value: 90},
		{Symbol: "L", Value: 50},
		{Symbol: "XL", Value: 40},
		{Symbol: "X", Value: 10},
		{Symbol: "IX", Value: 9},
		{Symbol: "V", Value: 5},
		{Symbol: "IV", Value: 4},
		{Symbol: "I", Value: 1},
	}

	sb := strings.Builder{}
	for i := 0; i < 13 && num != 0; i++ {
		romanSymbol := romanSymbols[i]
		for count := num / romanSymbol.Value; count != 0; count-- {
			sb.WriteString(romanSymbol.Symbol)
		}
		num = num % romanSymbol.Value
	}
	return sb.String()
}
