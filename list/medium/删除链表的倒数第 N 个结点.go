package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next:head}

	back, front := dummy, head
	for i := 0; i < n; i++ {
		if front == nil {
			return dummy.Next
		}
		front = front.Next
	}

	for front != nil {
		front = front.Next
		back = back.Next
	}

	back.Next = back.Next.Next

	return dummy.Next
}
