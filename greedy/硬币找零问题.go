package greedy

func getMinCoinCountHelper(total int, values []int, valueCount int) int {
	rest := total
	count := 0

	// 从大到小遍历所有面值
	for i := 0; i < valueCount; i++ {
		currentCount := rest / values[i] // 计算当前面值最多能用多少个
		rest %= values[i]  // 计算使用完当前面值后的余额
		count += currentCount

		if rest == 0 {
			return count
		}
	}

	return -1 // 无法凑出总价，返回-1
}


func GetMinCoinCount() int {
	values := []int{5, 3} // 硬币面值
	total := 11 // 总价

	return getMinCoinCountHelper(total, values, 2)
}
