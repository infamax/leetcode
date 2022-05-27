func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m, m)
	}
    
    if obstacleGrid[0][0] == 0 {
	    dp[0][0] = 1
    }

	for i := 1; i < n; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = dp[i - 1][0]
	}

	for j := 1; j < m; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = dp[0][j - 1]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
				dp[i][j] = 0
				continue
			}

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if obstacleGrid[i-1][j] == 1 {
				dp[i][j] += dp[i][j-1]
				continue
			}

			if obstacleGrid[i][j-1] == 1 {
				dp[i][j] += dp[i-1][j]
				continue
			}
			
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
    fmt.Println(dp)
	return dp[n-1][m-1]
}
