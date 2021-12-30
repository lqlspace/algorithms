package simple

// 删除第一个节点
func deleteFirstNode(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	pre := sentinel

	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
	}

	return sentinel.Next
}

// 删除所有值为val的节点
func deleteAllNode(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	pre := sentinel

	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
			continue
		}
		pre = pre.Next
	}

	return sentinel.Next
}
