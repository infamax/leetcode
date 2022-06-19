func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1, n+1)
	for i := 1; i*i <= n; i++ {
		dp[i*i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j*j < i; j++ {
			if dp[i] == 0 {
				dp[i] = dp[i-(j*j)] + 1
			} else {
				dp[i] = min(dp[i], dp[i-(j*j)] + 1)
			}
		}
	}

	return dp[n]
}
