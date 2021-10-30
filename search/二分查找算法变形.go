package search


//1. 查找第一个值等于给定值的元素
func BinarySearchFirst(arr []int, v int) int {
	low := 0
	high := len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == v {
			if mid == 0 || arr[mid-1] != v {
				return mid
			}
			high = mid-1
		} else if arr[mid] > v {
			high = mid-1
		} else {
			low = mid+1
		}
	}

	return -1
}
