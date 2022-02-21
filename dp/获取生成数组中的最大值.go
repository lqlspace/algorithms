package dp

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}

	nums := make([]int, n+1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + i%2*nums[i/2+1]
	}

	var ans int
	for _, v := range nums {
		ans = max(ans, v)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
