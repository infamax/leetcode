func runningSum(nums []int) []int {
	n := len(nums)
	prefixSum := make([]int, n, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i - 1] + nums[i]
	}
	return prefixSum
}
