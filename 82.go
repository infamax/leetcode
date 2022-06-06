func pushBack(tail *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = val
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var resHead, resTail *ListNode
	prev := head.Val
	head = head.Next
	count := 1
	for head != nil {
		if prev != head.Val {
			if count == 1 {
				if resHead == nil {
					resHead = new(ListNode)
					resHead.Val = prev
					resHead.Next = nil
					resTail = resHead
				} else {
					resTail = pushBack(resTail, prev)
				}
			}
			if head.Next == nil {
				if resHead == nil {
					resHead = new(ListNode)
					resHead.Val = head.Val
					resHead.Next = nil
					resTail = resHead
				} else {
					resTail = pushBack(resTail, head.Val)
				}
			}
			count = 1
		} else {
			count++
		}
		prev = head.Val
		head = head.Next
	}
	return resHead
}
