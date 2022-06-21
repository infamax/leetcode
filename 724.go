func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	prefixSum := make([]int, len(nums)+1, len(nums)+1)
	prefixSum[0] = 0
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	for i := 0; i < len(prefixSum) - 1; i++ {
		if prefixSum[i] == prefixSum[len(prefixSum)-1]-prefixSum[i+1] {
			return i
		}
	}
	return -1
}
