package backtrack

import (
	"sort"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。candidates 中的数字可以无限制重复被选取。
说明：所有数字（包括 target）都是正整数。解集不能包含重复的组合。
示例 1：
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[ [7], [2,2,3]]
示例 2：
输入：candidates = [2,3,5], target = 8,
所求解集为：
[ [2,2,2,2],[2,3,3], [3,5]]
 */

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var temp []int
	var dfs func(int, int)

	dfs = func(target, index int) {
		if target  == 0 {
			res = append(res, append([]int(nil), temp...))
			return
		}

		if target - candidates[index] < 0 {
			return
		}

		// for循环代表同一阶段的不同选择，backtrack函数代表下一阶段，而backtrack的第二个参数决定了下一个阶段的选择范围；
		for i := index; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(target - candidates[i], i)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(target, 0)

	return res
}
