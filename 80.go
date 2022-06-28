func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	i := 0
	for i < len(nums) {
		j := i
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}

		if j-i <= 2 {
            for k := 0; k < j - i; k++ {
                nums[index] = nums[i]
                index++
            }
            //fmt.Println("i = ", i)
            //fmt.Println("ok")
			//index += j - i
		} else {
			nums[index] = nums[i]
			index++
			nums[index] = nums[i]
			index++
		}
		i = j
	}
	return index
}
