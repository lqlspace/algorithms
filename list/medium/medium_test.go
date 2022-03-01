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
