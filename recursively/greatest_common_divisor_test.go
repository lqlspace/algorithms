/*
*执行性能测试，结果如下：执行五百万次，递归方法平均每次执行时间为303纳秒，循环迭代方法平均每次执行时间为277纳秒；
*ubuntu@VM-0-2-ubuntu:~/workspace/algorithm/src/algorithm-data-structure/algorithms/recursively$ go test greatest_common_divisor_test.go greatest_common_divisor.go -test.bench=".*"
*goos: linux
*goarch: amd64
*BenchmarkGCDRecur        5000000               303 ns/op
*BenchmarkGCDIteration    5000000               277 ns/op
*
*/
package recursive

import (
	"testing"
)

func TestGCDRecur(t *testing.T) {
	res := gcdRecur(8, 4)
	t.Logf("gcd of 8, 4 = %v", res)

	res = gcdRecur(4, 8)
	t.Logf("gcd of 4, 8 = %v", res)

	res = gcdRecur(0, 1)
	t.Logf("gcd of 0, 1 = %v", res)

	res = gcdRecur(0, 0)
	t.Logf("gcd of 0, 0 = %v", res)

	res = gcdRecur(7, 9)
	t.Logf("gcd of 7, 9 = %v", res)

	res = gcdRecur(6, 9)
	t.Logf("gcd of 6, 9 = %v", res)
}

func TestGCDIteration(t *testing.T) {
	res := gcdIteration(8, 4)
	t.Logf("gcd of 8, 4 = %v", res)

	res = gcdIteration(4, 8)
	t.Logf("gcd of 4, 8 = %v", res)

	
	res = gcdIteration(0, 1)
	t.Logf("gcd of 0, 1 = %v", res)

	res = gcdIteration(0, 0)
	t.Logf("gcd of 0, 0 = %v", res)

	res = gcdIteration(7, 9)
	t.Logf("gcd of 7, 9 = %v", res)

	res = gcdIteration(6, 9)
	t.Logf("gcd of 6, 9 = %v", res)
}


func BenchmarkGCDRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcdRecur(11203333333, 23478901234)
	}
}

func BenchmarkGCDIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcdIteration(11203333333, 23478901234)
	}
}

