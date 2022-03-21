package medium

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var combines []int
	var dfsCombination func(target, idx int)
	dfsCombination = func(target, idx int) {
		if idx == len(candidates) {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), combines...))
			return
		}

		dfsCombination(target, idx+1)

		if target - candidates[idx] >= 0 {
			combines = append(combines, candidates[idx])
			dfsCombination(target-candidates[idx], idx)
			combines = combines[:len(combines)-1]
		}
	}

	dfsCombination(target, 0)

	return ans
}
