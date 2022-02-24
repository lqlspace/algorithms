package advanced

func reverseKGroup(head *ListNode, k int) *ListNode {
	sentinel := &ListNode{Next: head}
	pre, cur := sentinel, head

	for cur != nil {
		for i := 0; i < k; i++ {
			if cur == nil {
				return sentinel.Next
			}
			cur = cur.Next
		}

		pre = reverseList(head, pre, cur)
		pre.Next = cur
		head = cur
	}

	return sentinel.Next
}

func reverseList(head *ListNode, pre, stop *ListNode) *ListNode {
	p := head

	for p != stop {
		tmp := p
		p = p.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}

	return head
}
