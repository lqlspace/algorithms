package simple

func balancedStringSplit(s string) int {
	var n, m, res int
	for _, val := range s {
		if val == 'L' {
			n++
		} else {
			m++
		}

		if n == m {
			res++
			n = 0
			m = 0
		}
	}

	return res
}
