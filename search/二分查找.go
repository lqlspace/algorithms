package search

func BinarySearch(a []int, v int) int {
	low := 0
	high := len(a)-1

	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
