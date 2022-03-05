package medium

func reorderList(head *ListNode)  {
	mid := midNode(head)
	right := reverseList(mid.Next)
	mid.Next = nil

	mergeLists(head, right)
}

func midNode(head *ListNode) *ListNode {
	dummy := &ListNode{Next:head}
	fast, slow := dummy, dummy

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		node := head
		head = head.Next
		node.Next = dummy.Next
		dummy.Next = node
	}

	return dummy.Next
}

func mergeLists(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		node := l2
		l2 = l2.Next
		node.Next = l1.Next
		l1.Next = node
		l1 = node.Next
	}
}

