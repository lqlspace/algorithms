package advanced

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}


type DLinkedNode struct {
	key, val int
	pre, next *DLinkedNode
}



func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.pre = l.head

	return l
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	node.pre.next = node
	node.next.pre = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeFromTail() *DLinkedNode {
	node := this.tail.pre
	node.pre.next = node.next
	node.next.pre = node.pre

	return node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(node)

	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	node = &DLinkedNode{key: key, val: value}
	this.addToHead(node)
	this.cache[key] = node
	this.size++

	if this.size > this.capacity {
		n := this.removeFromTail()
		delete(this.cache, n.key)
		this.size--
	}
}
