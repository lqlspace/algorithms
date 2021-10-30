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

func TestBinarySearchFirst(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2, 3, 4, 4}
	idx := BinarySearchFirst(arr, 2)
	t.Log(idx)
}

func TestBinarySearchLast(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2, 3, 4, 4, 5}
	idx := BinarySearchLast(arr, 2)
	t.Log(idx)

	idx = BinarySearchLast(arr, 4)
	t.Log(idx)
}

func TestBinarySearchGE(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 5, 6}
	idx := BinarySearchGE(arr, 3)
	t.Log(idx)
}

func TestBinarySearchLE(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 4, 5, 6}
	idx := BinarySearchLE(arr, 4)
	t.Log(idx)
}
