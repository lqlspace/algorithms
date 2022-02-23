package medium

func lengthOfLongestSubstring(s string) int {
	ans, n := 0, len(s)
	m := map[byte]struct{}{}

	for left, right := 0, 0; left < n; left++ {
		if left > 0 {
			delete(m, s[left-1])
		}

		for right < n {
			if _, ok := m[s[right]]; ok {
				break
			}
			m[s[right]] = struct{}{}
			right++
			ans = max(ans, len(m))
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


