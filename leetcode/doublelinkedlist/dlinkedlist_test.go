package doublelinkedlist

import (
	"testing"
)

func Test_LRUCache(t *testing.T)  {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 3)
	lruCache.Put(4, 1)
	v := lruCache.Get(1)
	t.Log(v)
	v = lruCache.Get(2)
	t.Log(v)
}
