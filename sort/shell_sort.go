package sort

//希尔排序是对插入排序的改进，分组可缩小无序数组规模，而后合并成部分有序数组也可以发挥插入排序的长处

func ShellSort(nums []int) []int {
	shell(nums, 0, len(nums)-1)
	return nums
}

func shell(nums []int, low, high int) {
	n := high - low + 1

	for key := n/2; key > 0; key /= 2 {
		for i := low+key; i <= high; i++ {
			for j := i; j >= key && nums[j] < nums[j-key]; j -= key {
				nums[j], nums[j-key] = nums[j-key], nums[j]
			}
		}
	}
}
