package heap

// build a heap
func buildHeap(a []int, n int)  {
	for i := n/2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

func heapifyUpToDown(a []int, top, count int)  {
	for i := top; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i*2
		}
		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2+1
		}

		if maxIndex == i  {
			break
		}
		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}
}

func Sort(a []int, n int)  {
	buildHeap(a, n)

	for k :=n; k >= 1; k-- {
		a[1], a[k] = a[k], a[1]
		heapifyUpToDown(a, 1, k-1)
	}
}
