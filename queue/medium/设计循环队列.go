package medium

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
