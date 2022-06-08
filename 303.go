type NumArray struct {
	prefixSum []int
}	


func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums) + 1, len(nums) + 1)
	prefixSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i - 1] + nums[i - 1]
	}
	return NumArray{
		prefixSum: prefixSum,
	}
}


func (n *NumArray) SumRange(left int, right int) int {
	return n.prefixSum[right + 1] - n.prefixSum[left]
}
