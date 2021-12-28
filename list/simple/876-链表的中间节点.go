package simple


// 快慢指针法（时间复杂度O(N), 空间复杂度O(1)），节点个数为偶数时，取得是中间两个节点中的后者
func middleNodeMethod1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 单指针法（时间复杂度O(N):须两遍循环, 空间复杂度O(1)）
func middleNodeMethod2(head *ListNode) *ListNode {
	var num int
	cur := head
	for cur != nil {
		num++
		cur = cur.Next
	}
	mid := num / 2
	for i := 0; i < mid; i++ {
		head = head.Next
	}

	return head
}

// 数组法（时间复杂度O(N):一遍循环，空间复杂度O(N)）
func middleNodeMethod3(head *ListNode) *ListNode {
	var num int
	arr := make([]*ListNode, 0)
	for head != nil {
		num++
		arr = append(arr, head)
		head = head.Next
	}
	return arr[num/2]
}


