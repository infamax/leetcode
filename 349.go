func intersection(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	
	f := 0
	s := 0
	prev := -1
	res := make([]int, 0)
	
	for i := 0; i < len(nums1) + len(nums2); i++ {
		if f >= len(nums1) {
			break
		}
		
		if s >= len(nums2) {
			break
		}
		
		if nums1[f] < nums2[s] {
			f++
		} else if nums1[f] == nums2[s] {
			if nums1[f] != prev {
				res = append(res, nums1[f])
				prev = nums1[f]
			}
			s++
			f++
		} else {
			s++
		}
	}
	return res
}
