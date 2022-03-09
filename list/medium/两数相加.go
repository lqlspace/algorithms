package medium

/*
给你两个链表表示两个非负数。各位数字是以倒序方式存在每个结点上的，
每个结点保存一个数字。把这两个数相加并且以同样的形式返回一个表示和的链表。
 */

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
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


// 时间复杂度O(MAX(M, N))， 空间复杂度O（MAX(M, N)）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	last := sentinel

	var carry int
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		last.Next = node
		last = last.Next
	}

	if carry > 0 {
		node := new(ListNode)
		node.Val = carry
		last.Next = node
		last = last.Next
	}

	return sentinel.Next
}
