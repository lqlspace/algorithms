package sort

/**********************堆排序*****************************/
func HeapSort(arr []int) {
	n := len(arr)

	// 建堆
	for i := n/2-1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 排序
	for i := n-1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}


func heapify(arr []int, n, idx int) {
	max := idx
	left := 2 * idx + 1
	right := 2 * idx + 2

	if left < n && arr[left] > arr[max] {
		max = left
	}
	if right < n && arr[right] > arr[max] {
		max = right
	}

	if max != idx {
		arr[max], arr[idx] = arr[idx], arr[max]
		heapify(arr, n, max)
	}
}

