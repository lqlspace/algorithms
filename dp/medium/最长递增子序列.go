package medium


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	var max int
	for i := 1; i < n; i++ {
		var tmpMax int
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && tmpMax < dp[j]{
				tmpMax = dp[j]
			}
		}
		dp[i] = tmpMax + 1
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
