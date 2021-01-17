package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 方法1：第一遍倒序，第二遍遍历(时间复杂度O(N), 空间复杂度O(1))
func reversePrintMethod1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	sentinel := new(ListNode)
	for head != nil {
		p := head
		head = head.Next
		p.Next = sentinel.Next
		sentinel.Next = p
	}

	arr := make([]int, 0)
	p := sentinel.Next
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}

	return arr
}

// 第一遍算链表元素个数，第二遍赋值(时间复杂度O(n),空间复杂度O(1))
func reversePrintMethod2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var num int
	for head != nil {
		num++
	}
	arr := make([]int, num)
	for head != nil {
		num--
		arr[num] = head.Val
	}

	return arr
}


// 使用栈
