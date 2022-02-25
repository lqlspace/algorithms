package simple

import (
	"math"
)

// 方法二： 遍历一遍
func maxProfit(prices []int) int {
	minPrice, maxProfit :=  math.MaxInt64, 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price - minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}
