package solutions

func MaxProfitIII(prices []int) int {
	prefixProfits := make([]int, len(prices))

	profit := 0
	bid := 10000
	for i, price := range prices {
		profit = max(profit, price-bid)
		prefixProfits[i] = profit
		if price < bid {
			bid = price
		}
	}

	profit = 0
	offer := 0
	for i := len(prices) - 1; i >= 0; i-- {
		price := prices[i]
		profit = max(profit, offer-price+prefixProfits[i])
		if price > offer {
			offer = price
		}
	}

	return profit
}
