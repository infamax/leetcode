func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	prefixSum := map[int]int{0: 1}
	count := 0
	sum := 0
	for _, item := range nums {
		sum += item
		c, ok := prefixSum[sum - k]
		if ok {
			count += c
		}
		prefixSum[sum]++
	}
	return count
}
