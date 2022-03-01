package advanced


// 循环两两合并
func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists2(ans, lists[i])
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

func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = pre.Next
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = pre.Next
			l2 = l2.Next
		}
	}

	if l1  != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}

	return dummy.Next
}


// 分治思想
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, low, high int) *ListNode {
	if low == high {
		return lists[low]
	}

	mid := low + (high-low) >> 1
	return mergeTwoLists(merge(lists, low, mid), merge(lists, mid+1, high))
}


