func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	oneCount := 0
	zeroCount := 0
	lastZeroIndex := 0
	maxLength := 0

	for left < len(nums) && right < len(nums) {
		if nums[right] == 1 {
			oneCount++
			maxLength = max(maxLength, oneCount)
			right++
		} else {
			zeroCount++
			if zeroCount == 1 {
				lastZeroIndex = right
				right++
			} else {
				left = lastZeroIndex + 1
				right = left
				oneCount = 0
				zeroCount = 0
			}
		}
	}
	
	if maxLength == len(nums) {
		return maxLength - 1
	}
	return maxLength
}
