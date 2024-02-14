package solutions

func MaxProfitII(prices []int) int {
	profit := 0
	prev := 10001
	bid := prev
	for _, price := range prices {
		if price < prev {
			profit += prev - bid
			bid = price
		}
		prev = price
	}
	profit += prev - bid
	return profit
}
