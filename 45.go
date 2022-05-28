const inf = math.MaxInt - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = inf
	}

	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] != inf {
				if j+nums[j] >= i {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}
	return dp[n-1]
}
