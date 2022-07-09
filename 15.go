func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	res := make([][]int, 0)
	m := make(map[struct {
		x int
		y int
		z int
	}]struct{})

	sort.Ints(nums)
    //fmt.Println("nums = ", nums)

	for i := 0; i < len(nums); i++ {
		x := nums[i]
		l := 0
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r] == -x && l != i && r != i {
                //fmt.Println("l = ", l)
                //fmt.Println("x = ", x)
                //fmt.Println("r = ", r)
				_, ok := m[struct {
					x int
					y int
					z int
				}{x: nums[i], y: nums[l], z: nums[r]}]
				if !ok {
                    //fmt.Println("Ok")
					res = append(res, []int{nums[i], nums[l], nums[r]})
				}
				m[struct {
					x int
					y int
					z int
				}{x: nums[i], y: nums[l], z: nums[r]}] = struct{}{}
				m[struct {
					x int
					y int
					z int
				}{x: nums[i], y: nums[r], z: nums[l]}] = struct{}{}
				m[struct {
					x int
					y int
					z int
				}{x: nums[l], y: nums[i], z: nums[r]}] = struct{}{}
				m[struct {
					x int
					y int
					z int
				}{x: nums[l], y: nums[r], z: nums[i]}] = struct{}{}
				m[struct {
					x int
					y int
					z int
				}{x: nums[r], y: nums[l], z: nums[i]}] = struct{}{}
				m[struct {
					x int
					y int
					z int
				}{x: nums[r], y: nums[i], z: nums[l]}] = struct{}{}
				l++
				r--
			} else if nums[l]+nums[r] > -x {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
