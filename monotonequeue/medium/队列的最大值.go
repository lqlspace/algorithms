package medium

type MaxQueue struct {
	queue []int
	dequeue []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		queue:   []int{},
		dequeue: []int{},
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.dequeue) == 0 {
		return -1
	}

	return this.dequeue[0]
}


func (this *MaxQueue) Push_back(value int)  {
	for len(this.dequeue) > 0 && value > this.dequeue[len(this.dequeue)-1] {
		this.dequeue = this.dequeue[:len(this.dequeue)-1]
	}
	this.queue = append(this.queue, value)
	this.dequeue = append(this.dequeue, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	ans := this.queue[0]
	this.queue = this.dequeue[1:]
	if ans == this.dequeue[0] {
		this.dequeue = this.dequeue[1:]
	}

	return ans
}
