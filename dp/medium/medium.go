package medium

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



// 动态规划
func longestPalindrome2(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	r := make([][]bool, l)
	for i := 0; i < l; i++ {
		r[i] = make([]bool, l)
	}

	ml, si := 1, 0
	// 字串长度 sl = j - i + 1
	for sl := 2; sl <= l; sl++ {
		// 左边界 i
		for i := 0; i < l; i++ {
			// 右边界 j - i + 1 = sl
			j := i + sl -1
			if j >= l{
				break
			}
			if s[i] != s[j] {
				r[i][j] = false
			} else {
				if sl <= 3 {
					r[i][j] = true
				} else {
					r[i][j] = r[i+1][j-1]
				}
			}

			if r[i][j] && sl > ml {
				ml = sl
				si = i
			}
		}
	}
	return s[si:si+ml]
}


