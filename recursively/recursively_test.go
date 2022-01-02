package recursive

import (
	"testing"
)

// 0, 1, 1, 2, 3, 5
func TestFibonacci(t *testing.T) {
	v1 := Fibonacci(3)
	v2 := FibonacciIteration(3)
	t.Log(v1, v2)

	v1 = Fibonacci(5)
	v2 = FibonacciIteration(5)
	t.Log(v1, v2)
}
