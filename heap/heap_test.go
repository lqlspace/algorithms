package heap

import (
	"fmt"
	"testing"
)

func TestHeap_Sort(t *testing.T) {
	heap := NewHeapFromArr([]int{3, 2, 1, 4, 5})
	arr := heap.Sort()

	fmt.Println(arr)
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
