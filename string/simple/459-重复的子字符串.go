package simple

// 时间复杂度O(N^2), 空间复杂度O(1)
func repeatedSubstringPattern(s string) bool {
	n := len(s)

	for i := 1; i <= n / 2; i++ {
		if n % i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}

	return false
}

// 使用kmp算法（拼接两个s，并去除第一个和最后一个字符）
// 时间复杂度O(N), 空间复杂度O(N)
func repeatedSubstringPattern2(s string) bool {
	return kmp(s + s, s)
}


func kmp(haystack, needle string) bool {
	n, m := len(haystack), len(needle)

	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[j] != needle[i] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	for i, j := 1, 0; i < n-1; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return true
		}
	}

	return false
}
