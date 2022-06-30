func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinItem(nums []int) int {
	minItem := nums[0]
	for i := 1; i < len(nums); i++ {
		minItem = min(minItem, nums[i])
	}
	return minItem
}

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minItem := MinItem(nums)
	count := 0
	for _, v := range nums {
		count += v - minItem
	}
	return count
}
