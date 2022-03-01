package medium

// 哈希表法
func detectCycle(head *ListNode) *ListNode {
	ht := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := ht[head]; ok {
			return head
		}
		ht[head] = struct{}{}
		head = head.Next
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return nil
}
