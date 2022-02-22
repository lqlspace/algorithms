package sort

/*
内层循环：遍历无序集合，找到min值对应索引下标;
外层循环：指定内层循环求出的min所在的索引，并交换；
时间复杂度：O(n^2)
 */
func SelectionSort(nums []int) []int {
	selectSort(nums, 0, len(nums)-1)
	return nums
}

func selectSort(nums []int, low, high int) {
	for i := low; i < high; i++ {
		min := i
		for j := i+1; j <= high; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
