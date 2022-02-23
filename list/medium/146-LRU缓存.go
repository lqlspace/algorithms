package medium

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	size int
	capacity int
	m map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, val int
	prev, next *DLinkedNode
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size:     0,
		capacity: capacity,
		m:        make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	this.moveToHead(this.m[key])

	return this.m[key].val
}


func (this *LRUCache) Put(key int, value int)  {
	if elem, ok := this.m[key]; ok {
		elem.val = value
		this.moveToHead(elem)
		return
	}

	node := &DLinkedNode{key: key, val: value}
	this.m[key] = node
	this.addToHead(node)
	this.size++
	if this.size > this.capacity {
		tail := this.removeTail()
		delete(this.m, tail.key)
		this.size--
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	this.head.next.prev = node
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)

	return node
}


