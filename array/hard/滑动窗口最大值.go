package hard

// 1. 暴力法
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
