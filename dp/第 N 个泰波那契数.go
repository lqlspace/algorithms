package dp

func tribonacci(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 2; i <=n; i++ {
		p, q, r = q, r, p+q+r
	}

	return r
}
