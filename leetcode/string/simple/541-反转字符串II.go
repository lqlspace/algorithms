package simple

// 时间复杂度O(N), 空间复杂度O(N)
func reverseStr(s string, k int) string {
	n := len(s)
	arr := []byte(s)

	for i := 0; i < n; i += 2 * k {
		for j, k := i, min(i+k-1, n-1); j < k; j++{
			arr[j], arr[k] = arr[k], arr[j]
			k--
		}
	}

	return string(arr)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
