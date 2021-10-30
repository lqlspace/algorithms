package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 4, 5, 7}
	idx := BinarySearch(arr, 4)

	t.Log(idx)

	idx = BinarySearch(arr, 3)

	t.Log(idx)

	arr = []int{}
	idx = BinarySearch(arr, 4)
	t.Log(idx)


}
