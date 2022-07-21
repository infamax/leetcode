func makeSliceFromList(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func reverseSliceBeetweenLandR(nums []int, l, r int) {
	for i := 0; i < (r - l + 1) / 2; i++ {
		nums[l + i], nums[r - i] = nums[r - i], nums[l + i] 
	}
}

func pushBack(tail *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = val
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func makeListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = nums[0]
	tail := head
	for i := 1; i < len(nums); i++ {
		tail = pushBack(tail, nums[i])
	}
	return head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	nums := makeSliceFromList(head)
	reverseSliceBeetweenLandR(nums, left - 1, right - 1)
	return makeListFromSlice(nums)
}
