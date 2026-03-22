func maxProfit(prices []int) int {
	r := 0
	result := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[r] {
			r = i
		}
		result = max(result, prices[i]-prices[r])
	}
	return result
}