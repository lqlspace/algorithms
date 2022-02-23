package medium

func lengthOfLongestSubstring(s string) int {
	ans, n := 0, len(s)
	m := make(map[byte]struct{})

	for i, j := 0, 0; i < n; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}

		for j < n {
			if _, ok := m[s[j]]; ok {
				break
			}

			m[s[j]] = struct{}{}
			ans = max(ans, len(m))
			j++
		}
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return  a
}



