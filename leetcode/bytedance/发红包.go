package bytedance

import (
	"math/rand"
	"time"
)

func sendRedPackage(n, k int) []int {
	v := n * 100
	arr := make([]int, k)

	rand.Seed(time.Now().Unix())

	for i := 0; i < v; i++ {
		arr[rand.Intn(k)]++
	}

	return arr
}
