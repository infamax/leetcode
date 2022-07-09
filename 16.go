func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) <= 2 {
		return -1
	}

	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		x := nums[i]
		l := 0
		r := len(nums) - 1
		for l < r {
			if l != i && r != i && abs(nums[l]+nums[r]+nums[x]-target) < abs(res-target) {
				res = nums[l] + nums[i] + nums[r]
			} else if nums[i]+nums[l]+nums[r] < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
