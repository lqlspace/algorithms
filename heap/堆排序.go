package heap

func Sort(arr []int, n int) {
	buildHeap(arr, n)

	for k := n; k >=1; k-- {
		arr[1], arr[k] = arr[k], arr[1]
		heapifyUpToDown(arr, 1, k-1)
	}
}

func buildHeap(a []int, n int) {
	for i := n/2; i >=1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

func heapifyUpToDown(a []int, top, count int) {
	for i := top; i <= count/2; {
		maxIndex := i
		if i * 2 <= count && a[maxIndex] < a[i*2] {
			maxIndex = i * 2
		}

		if i * 2 + 1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i * 2 + 1
		}
		if maxIndex == i {
			break
		}
		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}
}
