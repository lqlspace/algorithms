package sort

//插入排序对于部分有效的数组十分高效，也很适合小规模数组
func InsertSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1]  = nums[j-1], nums[j]
			}
		}
	}

	return nums
}
