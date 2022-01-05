package sort

/************************归并排序**************************************/

func MergeSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low+high) / 2
	sort(arr, low, mid)
	sort(arr, mid+1, high)
	merge(arr, low, mid, high)
}


func merge(arr []int, low, mid, high int) {
	merged := make([]int, 0, high-low+1)

	i, j := low, mid+1
	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			merged = append(merged, arr[i])
			i++
		} else {
			merged = append(merged, arr[j])
			j++
		}
	}
	if i <= mid {
		merged = append(merged, arr[i:mid+1]...)
	}
	if j <= high {
		merged = append(merged, arr[j:high+1]...)
	}

	for i := 0; i < len(merged); i++ {
		arr[low+i] = merged[i]
	}
}
