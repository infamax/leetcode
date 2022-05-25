func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var res int

	for k, v := range m {
		if v > len(nums) / 2 {
			res = k
			break
		}	
	}

	return res
}
