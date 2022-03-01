package advanced

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, right := dummy, head

	for right != nil {
		for i := 0; i < k; i++ {
			if right == nil {
				return dummy.Next
			}
			right = right.Next
		}
		pre = reverseList(pre, right)
	}

	return dummy.Next
}

func reverseList(pre, stop *ListNode) *ListNode {
	cur := pre.Next

	for cur.Next != stop {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return cur
}
