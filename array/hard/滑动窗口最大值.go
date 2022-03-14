package hard

import (
	"container/heap"
)

// 1. 优先队列
type Item struct {
	Val int
	Index int
}
type PriorityQueue struct {
	Items []*Item
}

func (pq *PriorityQueue) Len() int {
	return len(pq.Items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.Items[i].Val > pq.Items[j].Val
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	pq.Items = append(pq.Items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	item := pq.Items[len(pq.Items)-1]
	pq.Items = pq.Items[:len(pq.Items)-1]

	return item
}


func maxSlidingWindow(nums []int, k int) []int {
	pq := new(PriorityQueue)
	for i := 0; i < k; i++ {
		pq.Push(&Item{Val:nums[i], Index:i})
	}

	heap.Init(pq)
	ans := []int{pq.Items[0].Val}
	for i := k; i < len(nums); i++ {
		heap.Push(pq, &Item{Val:nums[i], Index:i})
		for pq.Items[0].Index < i - k + 1 {
			heap.Pop(pq)
		}

		ans = append(ans, pq.Items[0].Val)
	}

	return ans
}


// 2. 暴力法
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)

	var maxs []int
	for i := 0; i <= n - k; i++ {
		max := nums[i]
		for j := 1; j < k ; j++ {
			if nums[i+j] > max {
				max = nums[i+j]
			}
		}

		maxs = append(maxs, max)
	}

	return maxs
}
