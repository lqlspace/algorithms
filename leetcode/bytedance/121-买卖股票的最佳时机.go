package bytedance

import (
	"math"
)

// 方法-：暴力法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	res := 0
	for i := 0; i < len(prices) - 1; i++ {
		diff := 0
		for j := i+1; j < len(prices); j++ {
			if prices[j] - prices[i] > diff {
				diff = prices[j] - prices[i]
			}
		}

		if diff > res {
			res = diff
		}
	}

	return res
}

// 方法二： 遍历一遍
func maxProfit2(prices []int) int {
	minPoint := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPoint {
			minPoint = prices[i]
		} else if prices[i] - minPoint > maxProfit {
			maxProfit = prices[i] - minPoint
		}
	}

	return maxProfit
}
