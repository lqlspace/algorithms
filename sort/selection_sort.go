package sort

/*
内层循环：遍历无序集合，找到min值对应索引下标;
外层循环：指定内层循环求出的min所在的索引，并交换；
时间复杂度：O(n^2)
 */
func SelectionSort(arr []int, low, high int) {
	for i := low; i < high; i++ {
		min := i
		for j := i+1; j <= high; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
