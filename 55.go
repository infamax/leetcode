func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n, n)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := i; j >= 0; j-- {
			if dp[j] == true {
				if j + nums[j] >= i {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n - 1]
}
