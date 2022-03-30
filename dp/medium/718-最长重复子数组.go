package medium

func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	dp :=  make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	ans := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}

			ans = max(ans, dp[i][j])
		}
	}

	return ans 
}