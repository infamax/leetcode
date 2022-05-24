func pushBack(tail *ListNode, value int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = value
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
	res := new(ListNode)
	res.Val = head.Val
	res.Next = nil
	tail := res
	prev := head.Val
	for head != nil {
		if prev != head.Val {
			tail = pushBack(tail, head.Val)
		}
		prev = head.Val
		head = head.Next
	}
	return res
}
