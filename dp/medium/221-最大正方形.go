package medium

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	maxSide := dp[0][0]
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = int(matrix[i-1][j-1] - '0')
			if dp[i][j] == 1 {
				dp[i][j] = minFromThree(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1 
				if maxSide < dp[i][j] {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

func minFromThree(a, b, c int) int {
	min := a 
	if min > b {
		min = b  
	}
	if min > c {
		min = c 
	}

	return min 
}