package medium

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	return subSearch(nums, 0, n-1, target)
}

func subSearch(nums  []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := low + (high - low) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[low] <= nums[mid] {
		if nums[low] <= target && target <= nums[mid] {
			return subSearch(nums, low, mid-1, target)
		}
		return subSearch(nums, mid+1, high, target)
	} else {
		if nums[mid] <= target && target <= nums[high] {
			return subSearch(nums, mid+1, high, target)
		}
		return subSearch(nums, low, mid-1, target)
	}
}
