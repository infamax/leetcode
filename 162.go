const minInf = math.MinInt + 1

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findPeakElementHelp(nums []int, l, r int) int {
	m := (l + r) / 2
	var x, y int
	if m-1 == -1 {
		x = minInf
	} else {
		x = nums[m-1]
	}

	if m+1 == len(nums) {
		y = minInf
	} else {
		y = nums[m+1]
	}

	if nums[m] > x && nums[m] > y {
		return m
	}

	if r-l < 2 {
		return minInf
	}
	
	return max(findPeakElementHelp(nums, l, m), findPeakElementHelp(nums, m, r))
}


func findPeakElement(nums []int) int {
	return findPeakElementHelp(nums, -1, len(nums))
}
