package simple


func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0],dp[1] = 0, 0
	p, q, r := 0, 0, 0
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = min(p + cost[i-2], q + cost[i-1])
	}

	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
