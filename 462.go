func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minMoves2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	res := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	
	cur := nums[len(nums) / 2]
	
	for _, val := range nums {
		res += abs(val - cur)
	}
	return res
}
