package advanced


func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i])
	}

	return ans
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}


