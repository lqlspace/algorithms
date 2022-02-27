package medium

func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int

	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			ans = append(ans, append([]int(nil), nums...))
			return
		}

		for i := pos; i < n; i++ {
			nums[i], nums[pos] = nums[pos], nums[i]
			dfs(pos+1)
			nums[i], nums[pos] = nums[pos], nums[i]
		}
	}

	dfs(0)

	return ans
}
