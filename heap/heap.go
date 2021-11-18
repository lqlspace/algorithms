package heap

type Heap struct  {
	arr []int
	cap int
	count int
}

// init  heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		cap: capacity,
		arr: make([]int, capacity+1),
		count: 0,
	}
}

// 大顶堆 从下往上堆化
func (h *Heap) insert(data int) {
	// defensive
	if h.count == h.cap {
		return
	}

	h.count++
	h.arr[h.count] = data

	// compare with parent node
	i := h.count
	parent :=  i / 2
	for parent > 0 && h.arr[parent] < h.arr[i] {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i := parent
		parent =  i / 2
	}
}

// 删除堆顶
func (h *Heap) removeMax() {
	// defensive
	if h.count == 0 {
		return
	}

	h.arr[1], h.arr[h.count] = h.arr[h.count], h.arr[1]
	h.count--
}

// 从上往下堆化
func (h *Heap) heapifyUpToDown(a []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i * 2 + 1
		}

		if maxIndex == i {
			break
		}
		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}
}

