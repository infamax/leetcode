func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m, m)
	}
	
	dp[0][0] = grid[0][0]
	
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}
	
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j - 1] + grid[0][j]
	}
	
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]	
		}
	}
	
	return dp[n - 1][m - 1]
}
