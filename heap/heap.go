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

	i := h.count
	for parent := i/2; parent >= 1 && h.arr[parent] < h.arr[i]; parent = i / 2 {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
	}

	return nil
}

func (h *Heap) removeMax() error {
	if h.count == 0 {
		return errors.New("heap empty")
	}

	h.arr[1], h.arr[h.count] = h.arr[h.count], h.arr[1]
	h.count--
	h.heapifyUpToDown()

	return nil
}

func (h *Heap) heapifyUpToDown() {
	for i := 1; i <= h.count / 2; i++ {
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
