package medium

/*
给你两个链表表示两个非负数。各位数字是以倒序方式存在每个结点上的，
每个结点保存一个数字。把这两个数相加并且以同样的形式返回一个表示和的链表。
 */

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	dummy := new(ListNode)
	pre := dummy

	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node := &ListNode{Val: carry % 10}
		pre.Next = node
		pre = pre.Next
		carry /= 10
	}

	return dummy.Next
}
