package dp

// Best Time to Buy and Sell Stock
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/572/
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}
