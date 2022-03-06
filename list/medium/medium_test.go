package medium

import (
	"testing"
)

func TestLRUCache(t *testing.T)  {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	t.Log(lru.Get(1))
	lru.Put(3, 3)
	t.Log(lru.Get(2))
	lru.Put(4, 4)
	t.Log(lru.Get(1))
	t.Log(lru.Get(3))
	t.Log(lru.Get(4))
}

func Test_reverseBetween(t *testing.T)  {
	list := CreateList([]int{1, 2, 3, 4, 5})
	list = reverseBetween(list, 1, 3)
	TraverseList(list)

	list = CreateList([]int{1, 2, 3, 4, 5})
	list = reverseBetween(list, 2, 4)
	TraverseList(list)
}

func Test_reorderList(t *testing.T)  {
	list := CreateList([]int{1, 2, 3, 4})
	reorderList(list)
	TraverseList(list)

	list = CreateList([]int{1, 2, 3, 4, 5})
	reorderList(list)
	TraverseList(list)
}

func Test_removeNthFromEnd(t *testing.T)  {
	head := CreateList([]int{1,2,3,4,5})
	head = removeNthFromEnd(head, 2)
	TraverseList(head)

	head = CreateList([]int{1})
	head = removeNthFromEnd(head, 1)
	TraverseList(head)

	head = CreateList([]int{1, 2})
	head = removeNthFromEnd(head, 1)
	TraverseList(head)
}

func Test_deleteDuplicates(t *testing.T)  {
	head := CreateList([]int{1, 1, 2, 3, 3, 4, 4, 5})
	head = deleteDuplicates(head)
	TraverseList(head)

	head = CreateList([]int{1,1,1,2,3})
	head = deleteDuplicates(head)
	TraverseList(head)
}

func Test_sortList2(t *testing.T)  {
	head := CreateList([]int{4, 2, 1, 3})
	TraverseList(sortList2(head))

	head = CreateList([]int{4, 2, 1, 3})
	TraverseList(sortList3(head))

	head = CreateList([]int{4, 2, 1, 1, 3})
	TraverseList(sortList2(head))
	head = CreateList([]int{4, 2, 1, 1, 3})
	TraverseList(sortList3(head))
}
