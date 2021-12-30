package simple

// 迭代法（时间复杂度O(n), 空间复杂度O(1)）
func reverseListMethod1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentinel := new(ListNode)
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = sentinel.Next
		sentinel.Next = tmp
	}

	return sentinel.Next
}

// 递归法（时间复杂度O(N), 空间复杂度O(N)）
func reverseListMethod2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := reverseListMethod2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return pre
}
