package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) []int {
	rand.Seed(time.Now().UnixNano())

	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int)  {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high)
	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)

}

func partition(arr []int, low, high int) int {
	pos := rand.Intn(high-low+1) + low
	arr[pos], arr[high] = arr[high], arr[pos]

	i, pivot:= low, high
	for j := low; j < high; j++ {
		if arr[j] < arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[pivot] = arr[pivot], arr[i]

	return i
}
