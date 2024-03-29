package advanced

// dp
type pair struct {
	leftMax int
	rightMax int
}
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	dp := make([]pair, n)
	dp[0].leftMax, dp[n-1].rightMax = height[0], height[n-1]
	for i := 1; i < n; i++ {
		dp[i].leftMax = max(dp[i-1].leftMax, height[i])
	}
	for i := n-2; i >= 0; i-- {
		dp[i].rightMax = max(dp[i+1].rightMax, height[i])
	}

	var ans int
	for i := 0; i < n; i++ {
		ans += min(dp[i].leftMax, dp[i].rightMax) - height[i]
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}


// 双指针
func trap2(height []int) int {
	var ans int
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax =  max(rightMax, height[right])
		if leftMax <= rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}

	return ans
}
