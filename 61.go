func makeArrayFromList(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func rotateArray(a []int, k int) []int {
	res := make([]int, len(a), len(a))
	n := len(a)
	k %= n
	for i := 0; i < k; i++ {
		res[i] = a[n-k+i]
	}

	for i := k; i < n; i++ {
		res[i] = a[i-k]
	}

	return res
}

func pushBack(tail *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = val
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func makeListFromArray(a []int) *ListNode {
	head := new(ListNode)
	head.Val = a[0]
	head.Next = nil
	tail := head
	for i := 1; i < len(a); i++ {
		tail = pushBack(tail, a[i])
	}
	return head
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
	a := makeArrayFromList(head)
	b := rotateArray(a, k)
	return makeListFromArray(b)
}
