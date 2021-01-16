package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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
