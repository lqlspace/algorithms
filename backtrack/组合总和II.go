package backtrack

import (
	"sort"
)

/*
给出一组候选数 c 和一个目标数 t ，找出候选数中加起来和等于 t 的所有组合。c 中的每个数字在一个组合中只能使用一次。
 */

func CombinationSumII(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var path []int
	var backtrack func(target, index int)

	backtrack = func(target, index int) {
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}

		if index == len(candidates) {
			return
		}

		if target - candidates[index] < 0 {
			return
		}

		// for循环代表同一阶段的不同选择，backtrack函数代表下一阶段，而backtrack的第二个参数决定了下一个阶段的选择范围；
		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(target, 0)

	return res
}
