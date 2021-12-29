package basic

func PushBack(l *ListNode) *ListNode {
	sentinel := new(ListNode)
	last := sentinel

	for l != nil {
		p := l
		l = l.Next
		p.Next = nil
		last.Next = p
		last = last.Next
	}

	return sentinel.Next
}
