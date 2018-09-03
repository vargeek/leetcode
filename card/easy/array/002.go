package easy

// 买卖股票的最佳时机 II
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/22/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {
	count := len(prices)
	if count == 0 {
		return 0
	}
	min := prices[0]
	sum := 0
	for i := 0; i < count-1; i++ {
		if prices[i+1] < prices[i] {
			sum += prices[i] - min
			min = prices[i+1]
		}
	}
	sum += prices[count-1] - min
	return sum
}
