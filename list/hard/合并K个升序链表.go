package hard

import (
	"container/heap"
)

// 循环两两合并
func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists2(ans, lists[i])
	}

	return ans
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = pre.Next
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = pre.Next
			l2 = l2.Next
		}
	}

	if l1  != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}

	return dummy.Next
}


// 分治思想
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, low, high int) *ListNode {
	if low == high {
		return lists[low]
	}

	mid := low + (high-low) >> 1
	return mergeTwoLists(merge(lists, low, mid), merge(lists, mid+1, high))
}


// 优先队列
type priorityQueue []*ListNode

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

func (p *priorityQueue) Pop() interface{} {
	n := len(*p)
	item := (*p)[n-1]
	*p = (*p)[:n-1]

	return item
}

func mergeKLists4(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := new(ListNode)
	pre := dummy
	var pq priorityQueue
	for i := range lists {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}

	heap.Init(&pq)
	for len(pq) > 0 {
		node := heap.Pop(&pq).(*ListNode)
		pre.Next = node
		pre = pre.Next
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}

	return dummy.Next
}
