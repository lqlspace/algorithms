package heap

import (
	"errors"
	"fmt"
)

type Heap struct {
	arr []int
	cap int
	count int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		arr:   make([]int, capacity+1),
		cap:   capacity,
		count: 0,
	}
}

func NewHeapFromArr(arr []int) *Heap {
	h := NewHeap(len(arr))
	for _, item := range arr {
		h.insert(item)
	}

	return h
}

func (h *Heap) Sort() []int {
	if h.count == 0 {
		return []int{}
	}

	for k := h.count; k >=1; k-- {
		h.removeMax()
	}

	return h.arr[1:]
}


func (h *Heap) insert(val int) error {
	if h.count == h.cap {
		return errors.New("heap full")
	}

	h.count++
	h.arr[h.count] = val

	h.heapifyDownToUp(h.count)

	return nil
}

func (h *Heap) removeMax() error {
	if h.count == 0 {
		return errors.New("heap empty")
	}

	h.arr[1], h.arr[h.count] = h.arr[h.count], h.arr[1]
	h.count--
	h.heapifyUpToDown(1)

	return nil
}


func (h *Heap) buildHeap() {
	for i := h.count/2; i >= 1; i-- {
		h.heapifyUpToDown(i)
	}
}

// idx位置的值破坏了堆，重新构建
func (h *Heap) heapifyDownToUp(idx int)  {
	for i, parent := idx, idx/2; parent >= 1 && h.arr[parent] < h.arr[i]; parent = i / 2 {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
	}
}

//
func (h *Heap) heapifyUpToDown(top int) {
	for i := top; i <= h.count / 2; {
		maxIdx := i
		if h.arr[maxIdx] < h.arr[i*2] {
			maxIdx = i * 2
		}

		if i * 2 + 1 <= h.count && h.arr[maxIdx] < h.arr[i*2+1] {
			maxIdx = i * 2 + 1
		}

		if maxIdx == i {
			break
		}

		h.arr[i], h.arr[maxIdx]  = h.arr[maxIdx], h.arr[i]
		i =  maxIdx
	}
}

func (h *Heap) PrintHeap() {
	for i := 1; i <= h.count; i++ {
		fmt.Printf("%d ", h.arr[i])
	}
	println()
}
