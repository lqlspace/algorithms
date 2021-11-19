package queue

// 队列节点
type Node struct {
	value int
	priority  int
}


// 优先级队列
type PQueue struct {
	heap []Node
	capacity int
	used int
}

func NewPriorityQueue(capacity int) *PQueue {
	return &PQueue{
		heap:     make([]Node, capacity+1),
		capacity: capacity,
		used:     0,
	}
}

func (q *PQueue) Push(node Node) {
	if q.used > q.capacity {
		return
	}

	q.used++
	q.heap[q.used] = node
	//adjustHeap(q.heap, 1, q.used)

}

func (q *PQueue) Pop() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	adjustHeap(q.heap, 1, q.used)
	node := q.heap[1]

	q.heap[1] = q.heap[q.used]
	q.used--

	return node
}

func adjustHeap(src []Node, start, end int) {
	if start >= end {
		return
	}

	for i := end/2; i >= start; i-- {
		max := i
		if  src[max].priority < src[i*2].priority {
			max = i * 2
		}
		if i*2+1 <= end && src[max].priority < src[i*2+1].priority {
			max = i * 2 + 1
		}
		if max == i {
			continue
		}
		src[max], src[i] = src[i], src[max]
	}
}

