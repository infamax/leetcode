func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxItemSlice(nums []int) int {
	maxItem := nums[0]
	for i := 1; i < len(nums); i++ {
		maxItem = max(maxItem, nums[i])
	}
	return maxItem
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	return maxItemSlice(dp)
}
