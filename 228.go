func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	last := nums[0]
	cur := nums[0]
	lastIndex := 0
	res = append(res, "")

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cur = nums[i]
		} else {
			if cur == last {
				res[lastIndex] = strconv.Itoa(last)
			} else {
				res[lastIndex] = strconv.Itoa(last) + "->" + strconv.Itoa(cur)
			}
			last = nums[i]
			cur = last
			lastIndex++
			res = append(res, "")
		}
	}

	if cur == last {
		res[lastIndex] = strconv.Itoa(last)
	} else {
		res[lastIndex] = strconv.Itoa(last) + "->" + strconv.Itoa(cur)
	}
	return res
}
