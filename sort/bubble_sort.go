package sort

// 时间复杂度O(N^2), 空间复杂度O(1)
func BubbleSort(nums []int) []int {
	n := len(nums)

	for i := n-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
