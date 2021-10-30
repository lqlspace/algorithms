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

func TestBinarySearchRecursive(t *testing.T) {
	arr := []int{1, 3, 4, 7}
	e  := 5

	idx := BinarySearchRecursive(arr, e)
	t.Log(idx)

	e =  4
	idx = BinarySearchRecursive(arr, e)
	t.Log(idx)

	arr =  []int{}
	idx = BinarySearchRecursive(arr, e)
	t.Log(idx)
}
