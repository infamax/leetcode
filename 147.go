func pushBack(tail *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Val = val
	tmp.Next = nil
	tail.Next = tmp
	return tmp
}

func pushAfter(cur *ListNode, val int) *ListNode {
	if cur.Next == nil {
		return pushBack(cur, val)
	}

	next := cur.Next
	tmp := new(ListNode)
	tmp.Val = val
	cur.Next = tmp
	tmp.Next = next
	return tmp
}

func insertItem(head *ListNode, tail *ListNode, val int) {
	for head != nil && head.Val < val {
		head = head.Next
	}

	if head != nil {
		oldValue := head.Val
		head.Val = val
		pushAfter(head, oldValue)
	} else {
		pushBack(tail, val)
	}
}

func isSortedList(head *ListNode) bool {
	for head.Next != nil {
		if head.Val > head.Next.Val {
			return false
		}
		head = head.Next
	}
	return true
}

func updateHead(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func insertionSortList(head *ListNode) *ListNode {
	if isSortedList(head) {
		return head
	}
	resHead := new(ListNode)
	resHead.Val = head.Val
	resHead.Next = nil
	resTail := resHead
	head = head.Next
	for head != nil {
		insertItem(resHead, resTail, head.Val)
		resTail = updateHead(resHead)
		head = head.Next
	}
	return resHead
}
