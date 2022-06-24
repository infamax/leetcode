
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		_, ok := m[head]
		if ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}
