package doublelinkedlist

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next *DLinkedNode
}


func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     new(DLinkedNode),
		tail:     new(DLinkedNode),
	}
	l.head.next = l.tail
	l.tail.pre = l.head

	return l
}


func (this *LRUCache) Get(key int) int {
	// 1. 从map查询
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// 2. 找到，则把该节点写入到头部(删除，插入)
	node.pre.next = node.next
	node.next.pre = node.pre

	node.pre = this.head
	node.next = this.head.next
	node.next.pre = node
	node.pre.next = node

	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	// 1. 写入map
	if node, ok := this.cache[key]; ok {
		node.value = value

		// 2. 添加到链表头部
		node.pre.next = node.next
		node.next.pre = node.pre
		node.pre = this.head
		node.next = this.head.next
		node.next.pre = node
		node.pre.next = node
		return
	}
	node := &DLinkedNode{key: key, value: value}
	this.cache[key] = node

	// 2. 添加到链表头部
	node.pre = this.head
	node.next = this.head.next
	node.next.pre = node
	node.pre.next = node

	// 3. size加1
	this.size++

	// 4. 判断size有没有超过capacity,超过则删除末尾节点
	if this.size > this.capacity {
		n := this.tail.pre
		n.pre.next = n.next
		n.next.pre = n.pre
		delete(this.cache, n.key)
		this.size--
	}
}
