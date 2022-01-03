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

		for i := index; i < len(candidates); i++ {
			if i-1 >= index && candidates[i] == candidates[i-1] {
				continue
			}
			if target - candidates[i] < 0 {
				return
			}

			path = append(path, candidates[i])
			backtrack(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(target, 0)

	return res
}
