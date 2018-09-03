package dp

// Coin Change
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/809/
// https://leetcode.com/problems/coin-change/

func coinChange1(coins []int, amount int) int {
	counts := map[int]int{0: 0}
	for _, coin := range coins {
		maxCount := amount / coin
		counts2 := map[int]int{}
		for count := 1; count <= maxCount; count++ {
			value := count * coin
			for v, c := range counts {
				newValue := value + v
				newCount := count + c
				if newValue > amount {
					continue
				}
				// fmt.Println(coin, "---->", c, v, count, value, newCount, newValue)
				if c2, ok := counts2[newValue]; ok {
					if newCount < c2 {
						counts2[newValue] = newCount
					}
				} else if cc, ok := counts[newValue]; !ok || newCount < cc {
					counts2[newValue] = newCount
				}
			}
		}
		// fmt.Println(counts2)
		for v, c := range counts2 {
			counts[v] = c
		}
		// fmt.Println(counts)
		// fmt.Println("------------------------")
	}

	if c, ok := counts[amount]; ok {
		return c
	}
	return -1
}

func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = max
	}
	dp[0] = 0
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}
	if dp[amount] < max {
		return dp[amount]
	}
	return -1
}
