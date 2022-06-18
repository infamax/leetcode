func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func longestOnes(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	onesCount := 0
	zeroCount := 0
	lastZeroIndex := 0
	maxLength := 0
	l, r := 0, 0
	for l < len(nums) && r < len(nums) {
		if nums[r] == 1 {
			onesCount++
			r++
			maxLength = max(maxLength, onesCount+zeroCount)
		} else {
			zeroCount++
			if zeroCount <= k {
				if zeroCount == 1 {
					lastZeroIndex = r
				}
				r++
			} else {
				lastZeroIndex++
				l = lastZeroIndex
				r = l
				zeroCount = 0
				onesCount = 0
			}
		}
	}
	maxLength = max(maxLength, zeroCount + onesCount)
	return maxLength
}
