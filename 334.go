const inf = math.MaxInt - 1

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := inf
	mid := inf

	for i := 0; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}

	return false
}
