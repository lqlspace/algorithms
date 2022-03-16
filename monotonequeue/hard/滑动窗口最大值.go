package hard

func maxSlidingWindow(nums []int, k int) []int {
	var monotoneQueue []int
	push := func(i int) {
		for len(monotoneQueue) > 0 && nums[i] >= nums[monotoneQueue[len(monotoneQueue)-1]] {
			monotoneQueue = monotoneQueue[:len(monotoneQueue)-1]
		}
		monotoneQueue = append(monotoneQueue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := []int{nums[monotoneQueue[0]]}
	n := len(nums)
	for i := k; i < n; i++ {
		push(i)
		for monotoneQueue[0] <= i - k {
			monotoneQueue = monotoneQueue[1:]
		}
		ans = append(ans, nums[monotoneQueue[0]])
	}

	return ans
}

