package medium


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归
func rotateRight1(head *ListNode, k int) *ListNode {
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



// 迭代
func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	num, cur :=1, head
	for cur.Next != nil {
		num++
		cur = cur.Next
	}
	last := cur

	step := k % num
	pos := num - (step+1)
	sentinel := &ListNode{Next:head}
	cur = sentinel.Next
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}
	if cur.Next == nil {
		return sentinel.Next
	}

	node := cur.Next
	cur.Next = nil
	last.Next = sentinel.Next
	sentinel.Next = node

	return sentinel.Next
}


1 2 3 4 5
5 1 2 3 4
4 5 1 2 3
3 4 5 1 2
2 3 4 5 1
1 2 3 4 5

5 1 2 3 4

6 % 5
