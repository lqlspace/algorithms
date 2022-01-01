package simple

// 递归
func RotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	num, cur :=0, head
	for cur != nil {
		num++
		cur = cur.Next
	}

	step := k % num

	sentinel := new(ListNode)

	var recurFunc func(cur *ListNode) int
	recurFunc = func(cur *ListNode) int {
		if cur == nil {
			return 0
		}
		num := recurFunc(cur.Next)
		if num == step {
			cur.Next = nil
		} else if num < step {
			cur.Next = sentinel.Next
			sentinel.Next = cur
		}
		num++

		return num
	}
	recurFunc(head)
	last := sentinel
	for last.Next != nil {
		last = last.Next
	}
	last.Next = head

	return sentinel.Next
}



// 双指针迭代
func RotateRight2(head *ListNode, k int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	former, latter:= sentinel, sentinel

	for i := 0; i < k; i++ {
		if former == nil {
			return sentinel.Next
		}
		former = former.Next
	}

	if former.Next == nil {
		return sentinel.Next
	}

	for former.Next != nil {
		former = former.Next
		latter = latter.Next
	}

	former.Next = sentinel.Next
	sentinel.Next = latter.Next
	latter.Next = nil

	return sentinel.Next
}

// 首尾相连成环
func RotateRight3(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	num, cur :=1,  head
	for cur.Next != nil {
		num++
		cur = cur.Next
	}
	cur.Next = head

	step := num - (k % num)

	for i := 0; i < step; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil

	return head
}
