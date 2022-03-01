package advanced

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T)  {
	t.Log(TraverseList(reverseKGroup(CreateList([]int{1,2,3,4,5}), 2)))
	t.Log(TraverseList(reverseKGroup(CreateList([]int{1,2,3,4,5}), 3)))
	t.Log(TraverseList(reverseKGroup(CreateList([]int{1,2,3,4,5}), 1)))
}

func Test_mergeKLists(t *testing.T)  {
	l1 := CreateList([]int{1, 4, 5})
	l2 := CreateList([]int{1, 3, 4})
	l3 := CreateList([]int{2, 6})
	lists := []*ListNode{l1, l2, l3}
	l := mergeKLists2(lists)
	t.Log(TraverseList(l))
}
