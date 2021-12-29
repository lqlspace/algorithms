package simple

// 删除第一个节点
func deleteFirstNode(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	pre := sentinel

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			head.Next = nil
			break
		}
		pre = head
		head = head.Next
	}

	return sentinel.Next
}

// 删除所有值为val的节点
func deleteAllNode(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	pre := sentinel

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			head = pre.Next
			continue
		}
		pre = head
		head = head.Next
	}

	return sentinel.Next
}
