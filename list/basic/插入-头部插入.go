package basic

func PushFront(l *ListNode) *ListNode {
	sentinel := new(ListNode)
	for l != nil {
		p := l
		l = l.Next
		p.Next = sentinel.Next
		sentinel.Next = p
	}

	return sentinel.Next
}
