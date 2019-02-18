package randnum

import (
	"time"
	"math/rand"
)

func RandomIntSlice(size int) []int {
	arr := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		//arr[i] = 30 + rand.Intn(70)  // 30 ~ 100
		arr[i] = rand.Intn(100)        // 0 ~ 100
	}
	return arr
}

