package simple

// 迭代法
const min  = -65534
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sentinel := &ListNode{Next:head}
	var dup = min
	cur := sentinel
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			dup = cur.Next.Val
			cur.Next = cur.Next.Next.Next
		} else {
			if cur.Next.Val == dup {
				cur.Next = cur.Next.Next
			} else {
				dup = min
				cur = cur.Next
			}
		}
	}

	if cur.Next != nil && cur.Next.Val == dup {
		cur.Next = cur.Next.Next
	}

	return sentinel.Next
}

// 递归法
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dup bool
	var recurFunc func(cur *ListNode) *ListNode
	recurFunc = func(cur *ListNode) *ListNode {
		if cur.Next == nil {
			return cur
		}

		next := recurFunc(cur.Next)
		var ret *ListNode
		switch {
		case cur.Val == next.Val:
			dup = true
			ret = next
		case cur.Val != next.Val && !dup:
			ret = cur
		case cur.Val != next.Val && dup:
			dup = false
			cur.Next = next.Next
			ret = cur
		}
		return ret
	}

	head = recurFunc(head)
	if dup {
		head = head.Next
	}

	return head
}

// 借助hash表
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	hashMap := make(map[int]int)
	cur := head
	for cur != nil {
		hashMap[cur.Val]++
		cur = cur.Next
	}

	sentinel := &ListNode{Next:head}
	cur = sentinel
	for cur.Next != nil {
		if val, ok := hashMap[cur.Next.Val]; ok && val > 1{
				cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return sentinel.Next
}
