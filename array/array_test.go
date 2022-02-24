package array

import (
	"testing"
)

func TestMerge(t *testing.T) {
	master := make([]int, 8)
	master[0] = 1
	master[1] = 3
	master[2] = 5
	sub := []int{2, 4, 6, 7, 8}

	Merge(master, 3, sub, 5)

	t.Log(master)
}

func TestFind(t *testing.T) {
	arr := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	found := Find(arr, 4, 4, 7)
	t.Log(found)

	found = Find(arr,4, 4, 5)
	t.Log(found)
}

func Test_findKthLargest(t *testing.T)  {
	t.Log(findKthLargest([]int{3,2,1,5,6,4}, 2))
	t.Log(findKthLargest2([]int{3,2,1,5,6,4}, 2))
}

func Benchmark_findKthLargest(b *testing.B)  {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest([]int{3,2,1,5,6,4}, 2)
	}
}

func Benchmark_findKthLargest2(b *testing.B)  {

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest2([]int{3,2,1,5,6,4}, 2)
	}
}


