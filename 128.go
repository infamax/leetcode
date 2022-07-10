func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLength := 1

	s := make(map[int]bool)

	for _, val := range nums {
		s[val] = true
	}

	for key := range s {
		length := 0
		if !s[key - 1] {
			for s[key] {
				length++
				key++
			}
			maxLength = max(maxLength, length)
		}
	}

	return maxLength
}
