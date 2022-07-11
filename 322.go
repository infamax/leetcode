func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const inf = math.MaxInt - 1

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1, amount+1)

	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}

	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i - coin] + 1)
			}
		}
	}

	if dp[amount] == inf {
        return -1
	}

	return dp[amount]
}
