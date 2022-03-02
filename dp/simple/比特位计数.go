package simple

import (
	"math/bits"
)

// 使用dp，时间复杂度O(N)
func countBits(n int) []int {
	bits := make([]int, n+1)
	for highBit, i := 0, 1; i <= n; i++ {
		if i & (i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}

	return bits
}

// 使用标准库
func countBits2(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = bits.OnesCount(uint(i))
	}

	return arr
}

// Brain Kernighan算法，时间复杂度O(nlogn)
func countBits3(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = countOneBits(i)
	}

	return arr
}

func countOneBits(x int) int {
	var ones int
	for x > 0 {
		x &= x-1
		ones++
	}

	return ones
}

// dp：最低有效位
func countBits4(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}

	return bits
}

func countBits5(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}

	return bits
}
