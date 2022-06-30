func pushBack(tail *ListNode, x int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = x
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddPointer := head
	evenPointer := head.Next
	resHead := new(ListNode)
	resHead.Val = oddPointer.Val
	resTail := resHead
	oddPointer = oddPointer.Next
	for oddPointer != nil {
		if oddPointer.Next == nil {
			break
		}
		oddPointer = oddPointer.Next
		resTail = pushBack(resTail, oddPointer.Val)
		oddPointer = oddPointer.Next
	}

	for evenPointer != nil {
		resTail = pushBack(resTail, evenPointer.Val)
		evenPointer = evenPointer.Next
		if evenPointer == nil {
			break
		}
		evenPointer = evenPointer.Next
	}
	return resHead
}
