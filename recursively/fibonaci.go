package recursive

/*
求解斐波那契数列的第n个数
		/ -- 0, n = 0
F(n) =  - -- 1, n = 1
		\ -- F(n-1) + F(n-2), n >= 2
 */

// 递归法
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 迭代法
func FibonacciIteration(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	pre2, pre1 := 0, 1
	for i := 2; i <= n; i++ {
		pre1, pre2 =  pre1 + pre2, pre1
	}

	return pre1
}
