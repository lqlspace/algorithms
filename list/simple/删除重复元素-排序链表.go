package simple

/*
重复节点相邻
 */

func deleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next ==  nil {
		return head
	}

	pre := head
	for pre.Next != nil {
		if pre.Val == pre.Next.Val {
			pre.Next = pre.Next.Next
			continue
		}
		pre = pre.Next
	}

	return head
}
