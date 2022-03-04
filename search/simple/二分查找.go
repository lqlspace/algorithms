package simple

// 递归
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	return binarySearch(nums, 0, n-1, target)

}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left) >> 1
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearch(nums, mid+1, right, target)
	} else {
		return binarySearch(nums, left, mid-1, target)
	}
}
