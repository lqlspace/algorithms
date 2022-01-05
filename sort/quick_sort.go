package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int)  {
	if low >= high {
		return
	}

	pivotIdx := partition(arr, low, high)
	quickSort(arr, low, pivotIdx-1)
	quickSort(arr, pivotIdx+1, high)

}

func partition(arr []int, low, high int) int {
	pivot := high

	i := low-1
	for j := low; j < high; j++ {
		if arr[j] < arr[pivot] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[pivot] = arr[pivot], arr[i+1]

	return i+1
}
