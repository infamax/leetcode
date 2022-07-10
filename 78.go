func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	n := len(nums)
	for x := 0; x < int(math.Pow(2, float64(n))); x++ {
		tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			if x & (1 << i) > 0 {
				tmp = append(tmp, nums[i])
			}
		}
		res = append(res, tmp)
	}
	return res
}
