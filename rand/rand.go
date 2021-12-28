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

func search(items []int, val int) bool {
	var exist bool
	for i := range items {
		if items[i] == val {
			exist = true
		}
	}
	return exist
}
func RandomDistinctIntSlice(size int) []int {
	arr := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())

	for len(arr) < size {
		tmp := rand.Intn(100)
		if exist := search(arr, tmp); !exist {
			arr = append(arr, tmp)
		}
	}
	return arr
}

