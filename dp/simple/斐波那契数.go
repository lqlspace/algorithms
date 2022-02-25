package simple

func fib(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}

func fib3(n int) int {
	if n < 2 {
		return n
	}

	p, q := 0, 1
	for i := 2;  i <= n; i++ {
		p, q = q, p + q
	}

	return q
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib2(n-1) + fib2(n-2)
}
