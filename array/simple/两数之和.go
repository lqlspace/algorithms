package simple

func twoSum(nums []int, target int) []int {
	m, n := make(map[int]int), len(nums)

	for i := 0; i < n; i++ {
		if v, ok := m[target- nums[i]]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}

	return []int{}
}
