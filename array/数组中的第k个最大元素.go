package array

import (
	"math/rand"
	"time"
)

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

