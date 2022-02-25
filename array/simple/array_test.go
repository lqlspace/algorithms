package simple

import (
	"testing"
)

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

func Test_twoSum(t *testing.T)  {
	t.Log(twoSum([]int{2,7,11,15}, 9))
	t.Log(twoSum([]int{3, 2, 4}, 6))
	t.Log(twoSum([]int{3, 3}, 6))
}

func Test_merge(t *testing.T)  {
	nums1, m, nums2, n := []int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3
	merge(nums1, m, nums2, n)
	t.Log(nums1)

	nums1, m, nums2, n = []int{1}, 1, []int{}, 0
	merge(nums1, m, nums2, n)
	t.Log(nums1)

	nums1, m, nums2, n = make([]int, 0, 1), 0, []int{1}, 1
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}

