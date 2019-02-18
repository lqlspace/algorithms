package main

import (
	"time"
	"math/rand"
	"fmt"
)

func generateSlice(size int) []int {
	arr := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = 30 + rand.Intn(70) 
	}
	return arr
}

func bubbleSort(items []int) {
	n := len(items)

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if items[j] < items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}


func main() {
	s1 := generateSlice(20)
	bubbleSort(s1)
	fmt.Println(s1)
	s2 := generateSlice(20)
	bubbleSort(s2)
	fmt.Println(s2)
}
