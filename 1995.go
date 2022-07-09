func countQuadruplets(nums []int) int {
	if len(nums) < 4 {
		return -1
	}

	count := 0
	
	for i := 0; i < len(nums) - 3; i++ {
		for j := i + 1; j < len(nums) - 2; j++ {
			for k := j + 1; k < len(nums) - 1; k++ {
				for q := k + 1; q < len(nums); q++ {
					if nums[i] + nums[j] + nums[k] == nums[q] {
						count++
					}
				}
			}
		}
	}
	return count
}
