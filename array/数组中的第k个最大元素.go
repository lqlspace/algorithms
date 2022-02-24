package array

import (
	"math/rand"
	"time"
)

// 快排查询：时间复杂度O(N), 空间复杂度O(NlogN)
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)- k)
}

func quickSelect(nums []int, low, high, target int) int {
	pivot := randomPartition(nums, low, high)
	if pivot == target {
		return nums[pivot]
	} else if pivot <  target {
		return quickSelect(nums, pivot+1, high, target)
	} else {
		return quickSelect(nums, low, pivot-1, target)
	}
}

func randomPartition(nums []int, low, high int) int {
	p := rand.Intn(high-low+1) + low
	nums[p], nums[high] = nums[high], nums[p]

	i := low
	for j := low; j < high; j++ {
		if nums[j] < nums[high] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]

	return i
}

func findKthLargest2(nums []int, k int) int {
	return heapSelect(nums, k)
}

func heapSelect(nums []int, k int) int {
	n := len(nums)
	for i := (n-1)/2; i >= 0; i-- {
		heapify(nums, n, i)
	}

	for i := n-1; i > n-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums[0]
}


func heapify(nums []int, n, idx int) {
	max := idx
	l := 2 * idx + 1
	r := 2 * idx + 2
	if l < n && nums[l] > nums[max] {
		max = l
	}
	if r < n && nums[r] > nums[max] {
		max = r
	}
	if max != idx {
		nums[max], nums[idx] = nums[idx], nums[max]
		heapify(nums, n, max)
	}
}

