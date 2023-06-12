package n121

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		p := prices[i]
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}
