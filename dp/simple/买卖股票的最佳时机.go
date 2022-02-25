package simple

import (
	"math"
)

// 方法二： 遍历一遍
func maxProfit(prices []int) int {
	minPrices := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		} else if prices[i] - minPrices > maxProfit {
			maxProfit = prices[i]
		}
	}

	return maxProfit
}
