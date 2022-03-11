package medium

func nextPermutation(nums []int)  {
	if len(nums) == 0 {
		return
	}

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums)-1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; left, right = left + 1, right - 1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

