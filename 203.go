func pushBack(tail *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = val
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	var resHead *ListNode
	var resTail *ListNode

	for head != nil {
		if head.Val != val {
			if resHead == nil {
				resHead = new(ListNode)
				resHead.Val = head.Val
				resHead.Next = nil
				resTail = resHead
			} else {
				resTail = pushBack(resTail, head.Val)
			}
		}
		head = head.Next
	}
	return resHead
}
