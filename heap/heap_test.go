package heap

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{0, 3, 2, 1, 4, 5}
	Sort(arr, 5)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

func TestNewHeap(t *testing.T) {
	heap := NewHeap(4)

	heap.insert(3)
	heap.insert(5)
	heap.insert(1)
	heap.insert(9)

	heap.removeMax()

	heap.PrintHeap()

	heap.insert(2)
	heap.removeMax()

	heap.PrintHeap()
}