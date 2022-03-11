package medium

import (
	"sync"
)

type MyCircularQueue struct {
	queue []int
	head int
	size int
	capacity int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:    make([]int, k),
		head:     0,
		size:     0,
		capacity: k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	tail := (this.head + this.size) % this.capacity
	this.queue[tail] = value
	this.size++

	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % this.capacity
	this.size--

	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.head]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	rear := (this.head + this.size - 1) % this.capacity
	return this.queue[rear]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}


// 线程安全
type MyCircularQueueSafely struct {
	queue []int
	head int
	size int
	capacity int
	lock sync.Mutex
}


func ConstructorSafely(k int) MyCircularQueueSafely {
	return MyCircularQueueSafely{
		queue:    make([]int, k),
		head:     0,
		size:     0,
		capacity: k,
	}
}


func (this *MyCircularQueueSafely) EnQueue(value int) bool {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.IsFull() {
		return false
	}

	tail := (this.head + this.size) % this.capacity
	this.queue[tail] = value
	this.size++

	return true
}


func (this *MyCircularQueueSafely) DeQueue() bool {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % this.capacity
	this.size--

	return true
}


func (this *MyCircularQueueSafely) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.head]
}


func (this *MyCircularQueueSafely) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	rear := (this.head + this.size - 1) % this.capacity
	return this.queue[rear]
}


func (this *MyCircularQueueSafely) IsEmpty() bool {
	return this.size == 0
}


func (this *MyCircularQueueSafely) IsFull() bool {
	return this.size == this.capacity
}



// 链表实现
type MyCircularQueueLinked struct {
	head, tail *LinkedList
	size int
	capacity int
}

type LinkedList struct {
	Val int
	Next *LinkedList
}


func ConstructorLinked(k int) MyCircularQueueLinked {
	var ans MyCircularQueueLinked
	r := new(LinkedList)
	p := r
	for i := 1; i < k; i++ {
		p.Next = new(LinkedList)
		p = p.Next
	}
	p.Next = r
	ans.head, ans.tail = r, r
	ans.capacity = k

	return ans
}


func (this *MyCircularQueueLinked) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tail.Val = value
	this.tail = this.tail.Next

	this.size++

	return true
}


func (this *MyCircularQueueLinked) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = this.head.Next
	this.size--

	return true
}


func (this *MyCircularQueueLinked) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.head.Val
}


func (this *MyCircularQueueLinked) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	p := this.head
	for p.Next != this.tail {
		p = p.Next
	}

	return p.Val
}


func (this *MyCircularQueueLinked) IsEmpty() bool {
	return this.size == 0
}


func (this *MyCircularQueueLinked) IsFull() bool {
	return this.size == this.capacity
}

