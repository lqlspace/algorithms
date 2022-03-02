package medium

// 中心扩展法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1 - l1 > end - start {
			start, end = l1, r1
		}
		if r2 - l2 > end - start {
			start, end = l2, r2
		}
	}

	return s[start:end+1]
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for l >= 0 && r <=len(s)-1 && s[l] == s[r] {
		l, r = l-1, r+1
	}

	return l+1, r-1
}



// 动态规划法
func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	arr := make([][]bool, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]bool, n)
		arr[i][i] = true
	}

	start, max := 0, 1
	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				break
			}

			if s[i] == s[j] {
				if j - i + 1 < 3 {
					arr[i][j] = true
				} else {
					arr[i][j] = arr[i+1][j-1]
				}
			}

			if arr[i][j] && j - i + 1 > max {
				max = j - i + 1
				start = i
			}
		}
	}

	return s[start:start+max]
}


