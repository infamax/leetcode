func isIncreaseMonotonic(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func isDecreaseMonotonic(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			return false
		}
	}
	return true
}

func isMonotonic(nums []int) bool {
    if len(nums) == 1 {
        return true
    }
	if nums[1] > nums[0] {
		return isIncreaseMonotonic(nums)
    } else if nums[1] == nums[0] {
        fl1 := isIncreaseMonotonic(nums)
        fl2 := isDecreaseMonotonic(nums)
        if fl1 || fl2 {
            return true
        }
        return false
    }
	return isDecreaseMonotonic(nums)
}
