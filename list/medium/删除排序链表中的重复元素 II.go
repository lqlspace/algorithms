package medium

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		if pre.Next.Val == pre.Next.Next.Val {
			val := pre.Next.Val
			for pre.Next != nil && pre.Next.Val == val {
				pre.Next = pre.Next.Next
			}
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}
