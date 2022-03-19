package simple

func findMaxAverage(nums []int, k int) float64 {
	var maxSum, sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
