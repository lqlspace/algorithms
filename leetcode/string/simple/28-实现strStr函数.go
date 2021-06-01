package simple

// 1、滑动窗口暴力法，时间复杂度O(n*m)
func strStr1(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

outer:
	for i := 0; i <= n - m; i++ {
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}

	return -1
}

// 2、KMP算法，时间复杂度O（n+m）
func strStr2(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

func strStr3(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	for i := 0; i <= n - m; i++ {
		flag := true
		for j := range needle {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
