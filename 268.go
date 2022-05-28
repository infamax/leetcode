func missingNumber(nums []int) int {
	n := len(nums)
	sum := (1 + n) * n / 2
	s := 0
	for _, v := range nums {
		s += v
	}
	return sum - s
}
