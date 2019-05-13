package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	datas := make([]int, 50)
	for i := 0; i < 50; i++ {
		datas[i] = rand.Intn(50)
	}

	data1 := make([]int, 50)
	copy(data1, datas)

	QuickSort(data1, 0, len(data1)-1)

	fmt.Println(data1)

	data2 := make([]int, 50)
	copy(data2, datas)

	HeapSort(data2)

	fmt.Println(data2)

	data3 := make([]int, 50)
	copy(data3, datas)

	MergeSort(data3)

	fmt.Println(data3)

	data4 := make([]int, 50)
	copy(data4, datas)

	ShellSort(data4, 0, len(data4)-1)

	fmt.Println(data4)


	data5 := make([]int, 50)
	copy(data5, datas)

	InsertionSort(data5, 0, len(data5)-1)
	fmt.Println(data5)

	data6 := make([]int, 50)
	copy(data6, datas)

	SelectionSort(data6, 0, len(data6)-1)
	fmt.Println(data6)

	data7 := make([]int, 50)
	copy(data7, datas)

	BubbleSort(data7, 0, len(data7)-1)
	fmt.Println(data7)

	data8 := []int{9, 3, 6, 14, 2, 7, 5, 8, 1, 0, 13, 27, 38, 29, 28, 48}
	

	BitmapSort(data8, 0, 50)
	fmt.Println(data8)
}


func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	idx := Partition(arr, low, high)
	QuickSort(arr, low, idx-1)
	QuickSort(arr, idx+1, high)
}

func Partition(arr []int, low, high int) int {
	pivot := high
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < arr[pivot] {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[i+1], arr[pivot] = arr[pivot], arr[i+1]

	return i + 1
}


/**********************堆排序*****************************/
func HeapSort(arr []int) {
	length := len(arr)

	for i := length/2-1; i >= 0; i-- {
		Heapify(arr, length, i)
	}

	for i := length-1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Heapify(arr, i, 0)
	}
}


func Heapify(arr []int, n, idx int) {
	largest := idx
	left := 2 * idx + 1
	right := 2 * idx + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != idx {
		arr[largest], arr[idx] = arr[idx], arr[largest]
		Heapify(arr, n, largest)
	}
}


/************************归并排序**************************************/

func MergeSort(arr []int) {
	aux := make([]int, len(arr))

	sort(aux, arr, 0, len(arr)-1)
}

func sort(aux, arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low+high) / 2
	sort(aux, arr, low, mid)
	sort(aux, arr, mid+1, high)
	merge(aux, arr, low, mid, high)
}

func merge(aux, arr []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = arr[i]
	}

	i := low 
	j := mid + 1

	for k := low; k <= high; k++{
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > high {
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			arr[k] = aux[j]
			j++
		} else {
			arr[k] = aux[i]
			i++
		}
	}


}


/***********************希尔排序****************************/
func ShellSort(arr []int, low, high int) {
	n := high - low + 1

	for key := n/2; key > 0; key /= 2 {
		for i := low + key; i <= high; i++ {
			for j := i; j >= key && arr[j] < arr[j-key]; j -= key {
				arr[j], arr[j-key] = arr[j-key], arr[j]
			}
		}
	}
}


/***********************插入排序*****************************/
func InsertionSort(arr []int, low, high int) {
	for i := low; i <= high; i++ {
		for j := i; j > low && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}


/********************选择排序*****************************/
func SelectionSort(arr []int, low, high int) {
	for i := low; i < high; i++ {
		min := i
		for j := i+1; j <= high; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}


/*********************冒泡排序****************************/
func BubbleSort(arr []int, low, high int) {
	for i := high; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

/************************bitmap*******************************/
func BitmapSort(arr []int, min, max int) {
	lens := (max-min) / 8 + 1
	s := make([]int, lens)

	for _, val := range arr {
		s[val/8] |= (1 << uint(val%8))
	}

	var k int
	for i := range s {
		for j := 0; j < 8; j++ {
			if s[i] & (1 << uint(j)) > 0 {
				arr[k] = i * 8 + j
				k++
			}
		}
	}
}
