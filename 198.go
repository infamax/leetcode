func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArray(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = nums[i]
		for j := i - 2; j >= 0; j-- {
			dp[i] = max(dp[i], dp[j] + nums[i])
		}
	}
	return maxArray(dp)
}
