package sort

import (
	"testing"

	"lqlspace/algorithms/random"
)

func TestBubbleSort(t *testing.T) {
	arr := random.RandomIntSlice(10, 0, 10)
	t.Log(arr)

	BubbleSort(arr)
	t.Log(arr)
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := random.RandomIntSlice(10000, 0, 10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func TestHeapSort(t *testing.T) {
	arr := random.RandomIntSlice(10, 1, 10)
	t.Log(arr)

	HeapSort(arr)
	t.Log(arr)
}

func BenchmarkHeapSort(b *testing.B) {
	arr := random.RandomIntSlice(10000, 0, 10000)

	for i := 0; i < b.N; i++ {
		HeapSort(arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := random.RandomIntSlice(10, 1, 10)
	t.Log(arr)

	InsertionSort(arr)
	t.Log(arr)
}

func BenchmarkInsertionSort(b *testing.B) {
	arr := random.RandomIntSlice(10000, 0, 10000)

	for i := 0; i < b.N; i++ {
		InsertionSort(arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := random.RandomIntSlice(10, 1, 10)
	t.Log(arr)

	MergeSort(arr)
	t.Log(arr)
}

func TestQuickStart(t *testing.T) {
	arr := random.RandomIntSlice(10, 1, 10)
	t.Log(arr)

	QuickSort(arr)
	t.Log(arr)
}

func Test_sortArray(t *testing.T) {
	t.Log(sortArray([]int{1, 3, 4, 2, 9, 7}))
	t.Log(sortArray([]int{9, 7, 5, 5, 3, 8, 7}))
	t.Log(sortArray([]int{}))
}
