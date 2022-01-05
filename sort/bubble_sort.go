package sort

// 时间复杂度O(N^2), 空间复杂度O(1)
func BubbleSort(arr []int) {
	n := len(arr)

	for i := n-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
