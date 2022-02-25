package simple


// 哈希表(时间复杂度O（N），空间复杂度O(N))
func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return false
}


// 快慢指针（时间复杂度O(N)，空间复杂度O(1)）
// 在环里，跑的快的最终一定会追上跑的慢的
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

