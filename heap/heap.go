package heap

import (
	"fmt"
)

type Heap struct {
	a []int
	cap int
	cnt int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		a:   make([]int, capacity+1),
		cap: capacity,
		cnt: 0,
	}
}

func (h *Heap) insert(data int) {
	if h.cnt == h.cap {
		return
	}

	h.cnt++
	h.a[h.cnt] = data

	i := h.cnt
	parent := i / 2
	for parent > 0 && h.a[parent] < h.a[i] {
		h.a[parent], h.a[i] = h.a[i], h.a[parent]
		i = parent
		parent = i / 2
	}
}

func (h *Heap) removeMax() {
	if h.cnt == 0 {
		return
	}

	h.a[1], h.a[h.cnt] = h.a[h.cnt], h.a[1]
	h.cnt--

	h.heapifyUpToDown()
}

func (h *Heap) heapifyUpToDown() {
	for i := 1; i <= h.cnt/2; {
		maxIdx := i
		if h.a[maxIdx] < h.a[i*2] {
			maxIdx = i * 2
		}

		if i*2+1 <= h.cnt && h.a[maxIdx] < h.a[i*2+1] {
			maxIdx = i * 2 + 1
		}
		if maxIdx == i {
			break
		}
		h.a[i], h.a[maxIdx] = h.a[maxIdx], h.a[i]
		i = maxIdx
	}
}

func (h *Heap) PrintHeap() {
	for i := 1; i <= h.cnt; i++ {
		fmt.Printf("%d  ", h.a[i])
	}
	println()
}
