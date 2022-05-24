func mergeInSlice(nums1 []int, m int, nums2 []int, n int) []int {
	res := make([]int, m+n, m+n)
	f := 0
	s := 0
	for i := 0; i < m+n; i++ {
		if f >= m {
			res[i] = nums2[s]
			s++
			continue
		}

		if s >= n {
			res[i] = nums1[f]
			f++
			continue
		}

		if nums1[f] < nums2[s] {
			res[i] = nums1[f]
			f++
		} else {
			res[i] = nums2[s]
			s++
		}
	}
	return res
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := mergeInSlice(nums1, m, nums2, n)
	for i := 0; i < m+n; i++ {
		nums1[i] = res[i]
	}
}
