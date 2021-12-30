package simple


// 递归（时间复杂度O(n+m)，空间复杂度O(n+m):递归函数需要消耗
//栈空间，栈空间的大小取决于递归调用的深度。因为该题中递归函数
//mergeTwoListsMethod1最多调用n+m次，故。。。）
func mergeTwoListsMethod1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsMethod1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsMethod1(l1, l2.Next)
		return l2
	}
}

// 迭代（时间复杂度O(n+m)，空间复杂度O(1):我们只需要常数的空间存放若干变量）
func mergeTwoListsMethod2(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	pre := sentinel

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}

	return sentinel.Next
}
