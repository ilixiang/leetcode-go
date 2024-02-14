package solutions

func MaxProfit(prices []int) int {
	profit := 0
	bid := 100001
	for _, price := range prices {
		profit = max(profit, price-bid)
		if price < bid {
			bid = price
		}
	}
	return profit
}
