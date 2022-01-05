package sort

//插入排序对于部分有效的数组十分高效，也很适合小规模数组
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
