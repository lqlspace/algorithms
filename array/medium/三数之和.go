package medium

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	var ans [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i] + nums[j] + nums[k] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[i] + nums[j] + nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return ans
}
