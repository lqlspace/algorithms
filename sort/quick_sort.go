package sort

func Partition(items []int, low, high int) int {
	pivot := items[high]
	i := low - 1

	for j := low; j < high; j++ {
		if items[j] < pivot {
			i++
			items[j], items[i] = items[i], items[j]
		}
	}

	items[i+1], items[high] = items[high], items[i+1]

	return i + 1

}


func QuickSort(items []int, low, high int) {
	if low >= high {
		return
	}
	
	pivotIdx := Partition(items, low, high)
	QuickSort(items, low, pivotIdx-1)
	QuickSort(items, pivotIdx+1, high)
}
