package simple

func containsNearbyDuplicate(nums []int, k int) bool {
	ht := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(ht, nums[i-k-1])
		}
		if _, exist := ht[nums[i]]; exist {
			return true
		}
		ht[nums[i]] = struct{}{}
	}

	return false
}
