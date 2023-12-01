package knapsack

// 给你一个可装载重量为W的背包和N个物品，每个物品有重量和价值两个属性。其中第i个
// 物品的重量为wt[i]，价值为val[i]，现在让你用这个背包装物品，最多能装的价值是
// 多少?

// 状态：
//   装载的物品个数在变化
//	 背包承受的重量在变化

// 选择：
//   包含第i个物品
//   不包含第i个物品

// 定义：前i个物品在容量为j的限制下，dp[i][j] 最多能装载的价值

// base case
// dp[0][...] = 0, dp[...][0] = 0

// wt [3, 2, 4, 5]
// val[4, 1, 3, 2]
// i == 3

func Knapsack(w, n int, wt, val []int) int {
	// base case初始化
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if (j - wt[i-1]) < 0 {
				// 这种情况下只能选择不装入背包
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j-wt[i-1]]+val[i-1], dp[i-1][j])
			}
		}
	}

	return dp[n][w]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
