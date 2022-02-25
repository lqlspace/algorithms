package simple

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	f := make([][26]int, m+1)
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j + 'a') {
				f[i][j] =  i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}

	add := 0
	for i := 0; i < n; i++ {
		if f[add][int(s[i]- 'a')] == m {
			return false
		}

		add = f[add][int(s[i]-'a')] + 1
	}

	return true
}




// 双指针,对于每个s，t都有比对一下
func isSubsequence2(s string, t  string) bool {
	i, j, m, n := 0, 0, len(s), len(t)
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == m
}
