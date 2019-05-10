package sorting 


func Heapify(items []int, n, idx int) {
	largest := idx
	left := 2 * idx + 1
	right := 2 * idx + 2

	if left < n && items[left] > items[largest] {
		largest = left
	}
	if right < n && items[right] > items[largest] {
		largest = right
	}

	if largest != idx {
		items[idx], items[largest] = items[largest], items[idx]
		Heapify(items, n, largest)
	}

}



func HeapSort(items []int) {
	length := len(items)
	//Build heap
	for i := length/2-1; i >= 0; i-- {
		Heapify(items, length, i)
	}

	for i := length - 1; i >= 0; i-- {
		items[0], items[i] = items[i], items[0]
		Heapify(items, i, 0)
	}
}

