package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {

	min := 1 << 32
	max := 1 >> 32
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else if max < prices[i]-min {
			max = prices[i] - min
		}
	}

	return max
}
