package greedy


func canJump(nums []int) bool {
	lastMost  := 0
	for i := 0; i < len(nums); i++ {
		if i > lastMost {
			return false
		}
		if i + nums[i] > lastMost {
			lastMost = i + nums[i]
		}
	}

	return true
}
