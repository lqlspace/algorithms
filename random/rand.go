package random

import (
	"time"
	"math/rand"
)


func RandomIntSlice(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = min + rand.Intn(max-min)
	}

	return arr
}

func RandomDistinctIntSlice(size, min, max int) []int {
	if size > max - min {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	m := make(map[int]struct{})

	arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		for {
			val := min + rand.Intn(max-min)
			if _, ok := m[val]; !ok {
				arr = append(arr, val)
				m[val] = struct{}{}
				break
			}
		}

	}

	return arr
}

