package simple

type MyQueue struct {
	stack []int
	swap []int
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	if this.Empty() {
		this.stack = append(this.stack)
	}

	for !this.Empty() {
		val := this.Pop()
		this.swap = append(this.swap, val)
	}

	this.stack = append(this.stack, x)
	for len(this.swap) > 0 {
		val := this.swap[len(this.swap)-1]
		this.swap = this.swap[:len(this.swap)-1]
		this.stack = append(this.stack, val)
	}
}


func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	return val
}


func (this *MyQueue) Peek() int {
	return this.stack[len(this.stack)-1]
}


func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
